package outfit

import (
	"math/rand"
	"strconv"

	"github.com/TobiasNickel/outfit_creator/internal/type/outfit"
	"github.com/TobiasNickel/outfit_creator/internal/utils/httpclient"
	"github.com/pkg/errors"
)

func GetRandomOutfit(gender, country string) (accessory, upper, under *outfit.APIResponseItem, err error) {
	accessory, err = GetRandomItem(gender, country, outfit.SECTION_ACCESSORIES)
	if err != nil {
		err = errors.Wrap(err, "could not get SECTION_ACCESSORIES")
		return
	}
	upper, err = GetRandomItem(gender, country, outfit.SECTION_UPPER)
	if err != nil {
		err = errors.Wrap(err, "could not get SECTION_UPPER")
		return
	}
	under, err = GetRandomItem(gender, country, outfit.SECTION_UNDER)
	if err != nil {
		err = errors.Wrap(err, "could not get SECTION_UNDER")
		return
	}
	return
}

func GetRandomItem(gender, country, section string) (*outfit.APIResponseItem, error) {
	count, err := GetCount(gender, country, section)
	if err != nil {
		return nil, errors.Wrap(err, "could not count")
	}
	if count == 0 {
		return nil, nil
	}
	randomNumber := rand.Int() % count
	item, err := getByOffset(gender, country, section, randomNumber)
	if err != nil {
		return nil, errors.Wrap(err, "could not count")
	}
	return item, nil
}

func GetCount(gender, country, section string) (int, error) {
	type Counter struct {
		TotalCount int `json:"totalCount"`
	}
	var counter Counter

	category, err := getCategory(gender, section)
	if err != nil {
		return 0, errors.Wrap(err, "could not get category")
	}

	httpClient := httpclient.New()
	err = httpClient.Get(
		"https://api.newyorker.de/csp/products/public/query"+
			"?filters[country]="+country+
			"&filters[gender]="+gender+
			"&filters[web_category]="+category+
			"&limit=1",
		&counter)
	if err != nil {
		return 0, errors.Wrap(err, "could not count")
	}

	return counter.TotalCount, nil
}
func getByOffset(gender, country, section string, offset int) (*outfit.APIResponseItem, error) {
	httpClient := httpclient.New()
	var response outfit.OutfitAPIResponse
	category, err := getCategory(gender, section)
	if err != nil {
		return nil, errors.Wrap(err, "could not get category")
	}

	err = httpClient.Get(
		"https://api.newyorker.de/csp/products/public/query"+
			"?filters[country]="+country+
			"&filters[gender]="+gender+
			"&filters[web_category]="+category+
			"&limit=1&offset="+strconv.FormatInt(int64(offset), 10),
		&response)
	if err != nil {
		return nil, errors.Wrap(err, "could not count")
	}

	return &response.Items[0], nil
}

func getCategory(gender, section string) (string, error) {
	if gender == outfit.GENDER_MALE {
		if section == outfit.SECTION_ACCESSORIES {
			return outfit.WEB_CATEGPRU_MALE_ACCESSORIES, nil
		} else if section == outfit.SECTION_UPPER {
			return outfit.WEB_CATEGPRU_MALE_UPPER_GARMENT, nil
		} else if section == outfit.SECTION_UNDER {
			return outfit.WEB_CATEGPRU_MALE_UNDER_GARMENT, nil
		} else {
			return "", errors.New("invalid section " + gender + " " + section)
		}
	} else if gender == outfit.GENDER_FEMALE {
		if section == outfit.SECTION_ACCESSORIES {
			return outfit.WEB_CATEGPRU_FEMALE_ACCESSORIES, nil
		} else if section == outfit.SECTION_UPPER {
			return outfit.WEB_CATEGPRU_FEMALE_UPPER_GARMENT, nil
		} else if section == outfit.SECTION_UNDER {
			return outfit.WEB_CATEGPRU_FEMALE_UNDER_GARMENT, nil
		} else {
			return "", errors.New("invalid section " + gender + " " + section)
		}
	} else {
		return "", errors.New("invalid gender " + gender)
	}
}
