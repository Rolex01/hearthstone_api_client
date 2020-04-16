package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rolex01/hearthstone_api_client/models"
	"strings"
)

func (c *Client) GetCard(idOrSlug string) (card *models.Card, err error) {
	var b []byte

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/hearthstone/cards/%s?locale=%s", idOrSlug, c.locale), "")
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &card)
	if err != nil {
		err = errors.New(err.Error() + ": len([]byte)" + string(len(b)))
	}

	return
}

func (c *Client) GetAllCards() (cards *models.Cards, err error) {
	var (
		b         []byte
		cardsPath *models.Cards
	)

	for i := 1; ; i++ {
		b, err = c.getURLBody(
			c.apiURL+fmt.Sprintf("/hearthstone/cards?locale=%s&page=%d&pageSize=10000",
				c.locale,
				i,
			), "")
		if err != nil {
			return
		}

		if err = json.Unmarshal(b, &cardsPath); err != nil {
			return
		}

		cards.Cards = append(cards.Cards, cardsPath.Cards...)
		cards.CardCount += cardsPath.CardCount
		if i == cardsPath.PageCount {
			cards.Page, cards.PageCount = 1, 1
			break
		}
	}

	return
}

func (c *Client) CardsSearchConstructed(selection, class, rarity, typeCard, minionType, keyword, textFilter, sort,
	order string, manaCost, attack, health, collectible []int, page, pageSize int) (cards *models.Cards, err error) {
	var (
		b         []byte
		cardsPath *models.Cards
	)

	url := c.apiURL + fmt.Sprintf("/hearthstone/cards?locale=%s&gameMode=constructed", c.locale)

	if selection != "" {
		url += fmt.Sprintf("&set=%s", selection)
	}

	if class != "" {
		url += fmt.Sprintf("&class=%s", class)
	}

	if rarity != "" {
		url += fmt.Sprintf("&rarity=%s", rarity)
	}

	if typeCard != "" {
		url += fmt.Sprintf("&type=%s", typeCard)
	}

	if minionType != "" {
		url += fmt.Sprintf("&minionType=%s", minionType)
	}

	if keyword != "" {
		url += fmt.Sprintf("&keyword=%s", keyword)
	}

	if textFilter != "" {
		url += fmt.Sprintf("&textFilter=%s", textFilter)
	}

	if manaCost != nil {
		url += fmt.Sprintf("&manaCost=%s", strings.Trim(strings.Replace(fmt.Sprint(manaCost), " ", ",", -1), "[]"))
	}

	if attack != nil {
		url += fmt.Sprintf("&attack=%s", strings.Trim(strings.Replace(fmt.Sprint(attack), " ", ",", -1), "[]"))
	}

	if health != nil {
		url += fmt.Sprintf("&health=%s", strings.Trim(strings.Replace(fmt.Sprint(health), " ", ",", -1), "[]"))
	}

	if collectible != nil {
		url += fmt.Sprintf("&collectible=%s", strings.Trim(strings.Replace(fmt.Sprint(collectible), " ", ",", -1), "[]"))
	}

	if sort != "" {
		url += fmt.Sprintf("&sort=%d", sort)
	}

	if order != "" {
		url += fmt.Sprintf("&order=%d", order)
	}

	if page != 0 {
		url = url + fmt.Sprintf("&page=%d", page)
	}

	if pageSize != 0 {
		url = url + fmt.Sprintf("&pageSize=%d", pageSize)
	}

	if page == 0 && pageSize == 0 {
		for i := 1; ; i++ {
			b, err = c.getURLBody(
				c.apiURL+fmt.Sprintf("%s&page=%d&pageSize=10000",
					url,
					c.locale,
					i,
				), "")
			if err != nil {
				return
			}

			if err = json.Unmarshal(b, &cardsPath); err != nil {
				return
			}

			cards.Cards = append(cards.Cards, cardsPath.Cards...)
			cards.CardCount += cardsPath.CardCount
			if i == cardsPath.PageCount {
				cards.Page, cards.PageCount = 1, 1
				break
			}
		}
	} else {
		b, err = c.getURLBody(url, "")
		if err != nil {
			return
		}

		err = json.Unmarshal(b, &cards)
	}

	return
}

func (c *Client) CardsSearchBattlegrounds(tier []string, minionType, keyword, textFilter, sort,
	order string, attack, health, page, pageSize int) (cards *models.Cards, err error) {
	return
}
