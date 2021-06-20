import requests
from json import loads
import datetime


def check_on_company(inn='1026605606620'):
    params = {'key': 'e806a74f1d4f2876dadf8b9e086c205e980fa704', 'req': inn}
    base_url = 'https://api-fns.ru/api/egr'

    resp = requests.get(base_url, params=params)

    data = resp.json()
    data = data['items']
    # print(data)
    if len(data) == 0:
        return False, 0, 0, 0

    data = data[0]['ЮЛ']

    ogrn = data['ОГРН']
    short_name = data['НаимСокрЮЛ']
    long_name = data['НаимПолнЮЛ']

    return True, short_name, long_name, ogrn


def check_fin_koefs(inn='6663003127'):
    params = {'key': 'abae0a98a7fab849440b13be7e898dc155955e0f', 'inn': inn}
    base_url = 'https://damia.ru/api-scoring/fincoefs'

    resp = requests.get(base_url, params=params)

    data = resp.json()
    # print(data)
    data = data[inn]

    if len(data) == 0:
        return 0

    scores_sum = 0
    for d in data.values():
        years = sorted(d.keys(), reverse=True)
        for y in years:
            if 'Балл' in d[y]:
                score = d[y]['Балл']
                scores_sum += score
                break

    return scores_sum / len(data)


def check_coart(inn='6663003127'):
    params = {'key': 'afc89ccbb1aa0b021b19fb6bd0fe02a55a6eb229',
              'q': inn,
              'role': '2',
              'from_date': (datetime.datetime.now() - datetime.timedelta(days=128)).strftime('%Y-%m-%d'),
              'to_date': datetime.date.today().strftime('%Y-%m-%d'),
              'exact': '1', }
    base_url = 'https://damia.ru/api-arb/dela'
    resp = requests.get(base_url, params=params)

    data = resp.json()
    # print(data)
    try:
        data = data['result']['Ответчик']

        return len(data)
    except:
        return 0


def check_scoring(inn='6663003127'):
    params = {'key': 'abae0a98a7fab849440b13be7e898dc155955e0f',
              'inn': inn,
              'model': '_problemCredit',
              }
    base_url = 'https://damia.ru/api-scoring/score'
    resp = requests.get(base_url, params=params)

    data = resp.json()
    # print(data)
    data = data[inn]['_problemCredit']

    if len(data) == 0:
        return [1 / 2, 5 / 2, 1 / 2]  # neutral

    data = data[max(data.keys())]

    return [data['РискЗнач'], data['БаллЗнач'], data['НадежностьЗнач']]


def analysis(avg_fin_k, lawsuits_cnt, smart_scores):
    score = 0

    if lawsuits_cnt < 2:
        score += 3
    elif lawsuits_cnt < 4:
        score += 1

    score += avg_fin_k

    score *= smart_scores[2]

    return score
