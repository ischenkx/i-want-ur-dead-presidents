package grabber

import (
	"encoding/json"
	"errors"
	dto "github.com/ischenkx/innotech-backend/services/grabbing/service/db/dto"
	models "github.com/ischenkx/innotech-backend/services/grabbing/service/db/models"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Grabber struct {
	FnsKey     string
	ScoringKey string
	ArbitrKey  string
}

func (g *Grabber) request(baseUrl string, params map[string]string) (map[string]interface{}, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	log.Printf("Query: %s", q.Encode())

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var mp map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&mp)
	if err != nil {
		return nil, err
	}

	log.Printf("Got response: %+v", mp)

	return mp, nil
}

//TODO scoring logic upgrade

func (g *Grabber) GrabNames(product models.Product) (dto.GrabNamesResponse, error) {
	//debug
	//_, _ = g.request("https://jsonplaceholder.typicode.com/posts/1", map[string]string{})

	mp, err := g.request("https://api-fns.ru/api/egr", map[string]string{
		"key": g.FnsKey,
		"req": product.Inn,
	})
	if err != nil {
		return dto.GrabNamesResponse{}, err
	}

	items, ok := mp["items"].([]map[string]interface{})
	if !ok {
		return dto.GrabNamesResponse{}, errors.New("wrong format")
	}
	if len(items) < 0 {
		return dto.GrabNamesResponse{}, errors.New("nothing found")
	}
	entity, ok := items[0]["ЮЛ"].(map[string]interface{})
	if !ok {
		return dto.GrabNamesResponse{}, errors.New("wrong format")
	}

	return dto.GrabNamesResponse{
		Name:     entity["НаимСокрЮЛ"].(string),
		FullName: entity["НаимПолнЮЛ"].(string),
	}, nil
}

func (g *Grabber) GrabFinKoefScore(product models.Product) (dto.GrabFinKoefScoreResponse, error) {
	mp, err := g.request("https://damia.ru/api-scoring/fincoefs", map[string]string{
		"key": g.ScoringKey,
		"inn": product.Inn,
	})
	if err != nil {
		return dto.GrabFinKoefScoreResponse{}, err
	}

	mp, ok := mp[product.Inn].(map[string]interface{})
	if !ok {
		return dto.GrabFinKoefScoreResponse{}, errors.New("wrong format")
	}

	score := int64(0)
	for _, v := range mp {
		//shows data until past year
		currentYear := time.Now().Year() - 2

		vMap, ok := v.(map[string]interface{})
		if !ok {
			return dto.GrabFinKoefScoreResponse{}, errors.New("wrong format")
		}

		vMap, ok = vMap[strconv.Itoa(currentYear)].(map[string]interface{})
		if !ok {
			return dto.GrabFinKoefScoreResponse{}, errors.New("wrong format")
		}

		s, ok := vMap["Балл"]
		if !ok {
			return dto.GrabFinKoefScoreResponse{}, errors.New("wrong format")
		}

		sInt, err := strconv.ParseInt(s.(string), 10, 64)
		if err != nil {
			return dto.GrabFinKoefScoreResponse{}, errors.New("wrong format")
		}

		score += sInt
	}

	return dto.GrabFinKoefScoreResponse{
		Score: int(score),
	}, nil
}

func (g *Grabber) GrabCourtScore(product models.Product) (dto.GrabCourtScoreResponse, error) {

	//scan last year
	timeNow := time.Now()
	timeBefore := timeNow.Add(- time.Hour * 24 * 365)

	mp, err := g.request("https://damia.ru/api-arb/dela", map[string]string{
		"key":       g.ArbitrKey,
		"inn":       product.Inn,
		"role":      "2",
		"from_date": timeBefore.Format("2006-01-02"),
		"to_date":   timeNow.Format("2006-01-02"),
		"exact":     "1",
	})
	if err != nil {
		return dto.GrabCourtScoreResponse{}, err
	}

	mp, ok := mp["result"].(map[string]interface{})
	if !ok {
		return dto.GrabCourtScoreResponse{}, errors.New("wrong format")
	}

	mp, ok = mp["Ответчик"].(map[string]interface{})
	if !ok {
		return dto.GrabCourtScoreResponse{}, errors.New("wrong format")
	}

	/*
	Суд: 0-1 привлечение:  3 балла
	2-3: 1 балл
	>3: 0 баллов
	 */

	trialsCnt := len(mp)
	if trialsCnt>3 {
		return dto.GrabCourtScoreResponse{
			Score: 0,
		}, err
	} else if trialsCnt < 2 {
		return dto.GrabCourtScoreResponse{
			Score: 1,
		}, err
	} else {
		return dto.GrabCourtScoreResponse{
			Score: 3,
		}, err
	}

}

func (g *Grabber) GrabSmartScore(product models.Product) (dto.GrabSmartScoreResponse, error) {
	mp, err := g.request("https://damia.ru/api-scoring/score", map[string]string{
		"key": g.ScoringKey,
		"inn": product.Inn,
	})
	if err != nil {
		return dto.GrabSmartScoreResponse{}, err
	}

	mp, ok := mp[product.Inn].(map[string]interface{})
	if !ok {
		return dto.GrabSmartScoreResponse{}, errors.New("wrong format")
	}

	mp, ok = mp["_problemCredit"].(map[string]interface{})
	if !ok {
		return dto.GrabSmartScoreResponse{}, errors.New("wrong format")
	}

	//shows data until past year
	currentYear := time.Now().Year() - 1

	mp, ok = mp[strconv.Itoa(currentYear)].(map[string]interface{})
	if !ok {
		return dto.GrabSmartScoreResponse{}, errors.New("wrong format")
	}

	risk, ok := mp["РискЗнач"]
	if !ok {
		return dto.GrabSmartScoreResponse{}, errors.New("wrong format")
	}
	score, ok := mp["БаллЗнач"]
	if !ok {
		return dto.GrabSmartScoreResponse{}, errors.New("wrong format")
	}
	reliability, ok := mp["НадежностьЗнач"]
	if !ok {
		return dto.GrabSmartScoreResponse{}, errors.New("wrong format")
	}

	//TODO may result to service fault :)
	_, _ = strconv.ParseFloat(risk.(string), 64)        //[0;1] the lower, the better
	scoreF, _ := strconv.ParseFloat(score.(string), 64) //[0;5] the lower, the better
	_, _ = strconv.ParseFloat(reliability.(string), 64) //[0;1] the higher, the better

	//TODO use floats in calculations
	return dto.GrabSmartScoreResponse{Score: int(math.Round(scoreF))}, nil

}
