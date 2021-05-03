package discounter

import (
	"encoding/csv"
	"goarch/app/domain/models"
	"goarch/app/domain/repositories"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const StartHour = 5
const requestUrl = "https://raw.githubusercontent.com/goarchitecture/lesson-2/feature/lesson-4/assets/discounts.csv"

type Discounter struct {
	ItemRepository repositories.ItemRepository
	CheckTime      time.Time
}

func (d *Discounter) download() (Discounts, error) {
	resp, err := http.DefaultClient.Get(requestUrl)
	if err != nil {
		return nil, err
	}

	lines, err := csv.NewReader(resp.Body).ReadAll()
	if err != nil {
		return Discounts{}, err
	}

	disc := make(Discounts, len(lines)-1)

	for i, line := range lines {
		if i == 0 {
			continue
		}
		discount, err := strconv.Atoi(line[2])
		if err != nil {
			continue
		}
		disc[i-1] = &Discount{
			Type:   DiscountType(line[0]),
			Value:  line[1],
			Amount: float64(discount),
		}
	}

	return disc, nil
}

func (d *Discounter) Start() error {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return err
	}

	now := time.Now().Local()
	firstCallTime := time.Date(
		now.Year(), now.Month(), now.Day(), StartHour, 0, 0, 0, loc)
	if firstCallTime.Before(now) {
		// Если получилось время раньше текущего, прибавляем сутки.
		firstCallTime = firstCallTime.Add(time.Hour * 24)
	}

	// Вычисляем временной промежуток до запуска.
	duration := firstCallTime.Sub(time.Now().Local())
	go func() {
		for {
			time.Sleep(duration)

			disc, err := d.download()
			if err != nil {
				continue
			}

			for _, dElem := range disc {
				var items []*models.Item
				switch dElem.Type {
				case Category:
					items, err = d.ItemRepository.GetWhereCategory(strings.ToLower(dElem.Value))
					if err != nil {
						continue
					}
				case Item:
					item, err := d.ItemRepository.Get(dElem.Value)
					if err != nil {
						continue
					}
					items = []*models.Item{item}
				case All:
					items, err = d.ItemRepository.GetAll()
					if err != nil {
						continue
					}
				default:
					continue
				}

				for _, item := range items {
					item.Price.Discounted = item.Price.Base * (100 - dElem.Amount) / 100
					if _, err := d.ItemRepository.Store(item); err != nil {
						continue
					}
				}
			}
		}
	}()

	return nil
}

func New(itemRepository repositories.ItemRepository) *Discounter {
	return &Discounter{ItemRepository: itemRepository}
}
