# 1. Počet odpovedí v ankete

Popis: Základná informácia – koľko ľudí vyplnilo dotazník.

```sql
SELECT COUNT(*) AS total_responses
FROM responses;
```

```
┌─────────────────┐
│ total_responses │
│      int64      │
├─────────────────┤
│       83        │
└─────────────────┘
```

---

# 2. Počet odpovedí podľa lokality

Popis: Zistíš, z ktorých častí mesta prišli respondenti.

```sql
SELECT
    option_label AS location,
    COUNT(*) AS responses,
    ROUND(100.0 * COUNT(*) / SUM(COUNT(*)) OVER (), 2) AS percentage
FROM response_options
WHERE question_code = 'location'
GROUP BY option_label
ORDER BY responses DESC;
```

```
┌─────────────────────────────────────────────────────────────────┬───────────┬────────────┐
│                            location                             │ responses │ percentage │
│                             varchar                             │   int64   │   double   │
├─────────────────────────────────────────────────────────────────┼───────────┼────────────┤
│ BB Sásová (Rudlová, Sásová)                                     │        49 │      57.65 │
│ BB Fončorda, Radvaň                                             │        19 │      22.35 │
│ BB Stred (Banská Bystrica-centrum)                              │        11 │      12.94 │
│ BB Juh (lliaš, Kráľová, Kremnička, Pršianska Terasa, Rakytovce) │         4 │       4.71 │
│ BB Východ (Majer, Senica, Šalková, Uhlisko)                     │         2 │       2.35 │
└─────────────────────────────────────────────────────────────────┴───────────┴────────────┘

Lokalita respondentov
BB Sásová                      49 | ######################################## 57.65%
BB Fončorda, Radvaň            19 | ###############                          22.35%
BB Stred                       11 | #########                                12.94%
BB Juh                          4 | ###                                       4.71%
BB Východ                       2 | ##                                        2.35%
```

---

# 3. Vek detí respondentov

Popis: Ukáže vekové kategórie detí rodičov v ankete.

```sql
SELECT
    option_label AS age_group,
    COUNT(*) AS responses,
    ROUND(100.0 * COUNT(*) / SUM(COUNT(*)) OVER (), 2) AS percentage
FROM response_options
WHERE question_code = 'age'
GROUP BY option_label
ORDER BY responses DESC;
```

```
┌────────────────────┬───────────┬────────────┐
│     age_group      │ responses │ percentage │
│      varchar       │   int64   │   double   │
├────────────────────┼───────────┼────────────┤
│ 1-3 rokov          │        37 │      33.04 │
│ 3-6 rokov          │        36 │      32.14 │
│ 6+ rokov           │        19 │      16.96 │
│ 0-1 rok            │        11 │       9.82 │
│ ešte sa nenarodilo │         9 │       8.04 │
└────────────────────┴───────────┴────────────┘

Vek dieťaťa
1-3 rokov                37 | #################################       33.04%
3-6 rokov                36 | ################################        32.14%
6+ rokov                 19 | #################                       16.96%
0-1 rok                  11 | ##########                               9.82%
ešte sa nenarodilo        9 | ########                                 8.04%
```

---

# 4. Preferovaný čas aktivít

Popis: Kedy chcú rodičia najčastejšie aktivity.

```sql
SELECT
    option_label AS preferred_time,
    COUNT(*) AS responses,
    ROUND(100.0 * COUNT(*) / SUM(COUNT(*)) OVER (), 2) AS percentage
FROM response_options
WHERE question_code = 'preffered_time'
GROUP BY option_label
ORDER BY responses DESC;
```

```
┌───────────────────────────────┬───────────┬────────────┐
│        preferred_time         │ responses │ percentage │
│            varchar            │   int64   │   double   │
├───────────────────────────────┼───────────┼────────────┤
│ podvečer (po 17:00)           │        37 │      27.82 │
│ popoludní (napr. 15:00–17:00) │        36 │      27.07 │
│ dopoludnia (napr. 9:00–12:00) │        36 │      27.07 │
│ víkendy                       │        24 │      18.05 │
└───────────────────────────────┴───────────┴────────────┘

Preferovaný čas aktivít
Podvečer                 37 | #################################       27.82%
Popoludní                36 | ################################        27.07%
Dopoludnia               36 | ################################        27.07%
Víkendy                  24 | #####################                   18.05%
```

---

# 5. Najžiadanejšie aktivity pre rodičov

Popis: Zistíš, ktoré typy aktivít majú rodičia najväčší záujem.

```sql
SELECT
    option_label AS activity,
    COUNT(*) AS responses,
    ROUND(100.0 * COUNT(*) / SUM(COUNT(*)) OVER (), 2) AS percentage
FROM response_options
WHERE question_code = 'activities_for_parent'
GROUP BY option_label
ORDER BY responses DESC;
```

```
┌──────────────────────────────────────────────────────────┬───────────┬────────────┐
│                         activity                         │ responses │ percentage │
│                         varchar                          │   int64   │   double   │
├──────────────────────────────────────────────────────────┼───────────┼────────────┤
│ prednášky / workshopy (výchova, psychológia, zdravie…)   │        45 │      23.56 │
│ cvičenie pre mamičky (jóga, fit po pôrode, pilates…)     │        35 │      18.32 │
│ tvorivé dielne                                           │        35 │      18.32 │
│ praktické kurzy (napr. varenie, prvá pomoc)              │        33 │      17.28 │
│ relaxačné aktivity (jóga, dychové cvičenia, mindfulness) │        22 │      11.52 │
│ podporné skupiny (diskusie, zdieľanie skúseností)        │        21 │      10.99 │
└──────────────────────────────────────────────────────────┴───────────┴────────────┘
Aktivity pre rodičov
Prednášky / workshopy    45 | ######################################## 23.56%
Cvičenie pre mamičky     35 | ###############################          18.32%
Tvorivé dielne           35 | ###############################          18.32%
Praktické kurzy          33 | #############################            17.28%
Relaxačné aktivity       22 | ###################                      11.52%
Podporné skupiny         21 | ##################                       10.99%
```

---

# 6. Najžiadanejšie aktivity pre deti

Popis: Najpopulárnejšie detské aktivity.

```sql
SELECT
    option_label AS activity,
    COUNT(*) AS responses,
    ROUND(100.0 * COUNT(*) / SUM(COUNT(*)) OVER (), 2) AS percentage
FROM response_options
WHERE question_code = 'activities_for_children'
GROUP BY option_label
ORDER BY responses DESC;
```

```
┌─────────────────────────────────────────┬───────────┬────────────┐
│                activity                 │ responses │ percentage │
│                 varchar                 │   int64   │   double   │
├─────────────────────────────────────────┼───────────┼────────────┤
│ pohybové aktivity (cvičenie, tanec)     │        54 │      23.08 │
│ herničky / voľná hra                    │        46 │      19.66 │
│ hudobné aktivity / rytmika              │        38 │      16.24 │
│ adaptačné programy (príprava na škôlku) │        34 │      14.53 │
│ jazykové aktivity                       │        31 │      13.25 │
│ výtvarné / kreatívne aktivity           │        31 │      13.25 │
└─────────────────────────────────────────┴───────────┴────────────┘
Aktivity pre deti
Pohybové aktivity        54 | ######################################## 23.08%
Herničky / voľná hra     46 | ##################################       19.66%
Hudobné / rytmika        38 | ############################              16.24%
Adaptačné programy       34 | #########################                 14.53%
Jazykové aktivity        31 | #######################                   13.25%
Výtvarné aktivity        31 | #######################                   13.25%
```

---

# 7. Prečo ľudia nenavštevujú centrum

Popis: Najčastejšie dôvody neúčasti.

```sql
SELECT
    option_label AS reason,
    COUNT(*) AS responses,
    ROUND(100.0 * COUNT(*) / SUM(COUNT(*)) OVER (), 2) AS percentage
FROM response_options
WHERE question_code = 'not_attending_reason'
GROUP BY option_label
ORDER BY responses DESC;
```

```
┌───────────────────────────┬───────────┬────────────┐
│          reason           │ responses │ percentage │
│          varchar          │   int64   │   double   │
├───────────────────────────┼───────────┼────────────┤
│ Kvôli vzdialenosti        │        15 │      34.88 │
│ Nenavštevujeme            │        12 │      27.91 │
│ Nevyhovuje mi čas aktivít │        11 │      25.58 │
│ Neatraktívne aktivity     │         3 │       6.98 │
│ Nepáčia sa mi priestory   │         2 │       4.65 │
└───────────────────────────┴───────────┴────────────┘
Dôvody neúčasti
Kvôli vzdialenosti        15 | ######################################## 34.88%
Nenavštevujeme            12 | ###############################          27.91%
Nevyhovuje čas            11 | ############################             25.58%
Neatraktívne aktivity      3 | #######                                   6.98%
Nepáčia sa priestory       2 | #####                                     4.65%
```

---

# 8. Najdôležitejšie faktory pri výbere aktivít

Popis: Čo je pre rodičov najdôležitejšie.

```sql
SELECT
    option_label AS factor,
    COUNT(*) AS responses,
    ROUND(100.0 * COUNT(*) / SUM(COUNT(*)) OVER (), 2) AS percentage
FROM response_options
WHERE question_code = 'preffered_cons'
GROUP BY option_label
ORDER BY responses DESC;
```

```
┌──────────────────────┬───────────┬────────────┐
│        factor        │ responses │ percentage │
│       varchar        │   int64   │   double   │
├──────────────────────┼───────────┼────────────┤
│ čas                  │        56 │      33.14 │
│ cena                 │        45 │      26.63 │
│ kvalita lektora      │        35 │      20.71 │
│ priateľská atmosféra │        33 │      19.53 │
└──────────────────────┴───────────┴────────────┘
Kľúčové faktory výberu
Čas                     56 | ######################################## 33.14%
Cena                    45 | ################################         26.63%
Kvalita lektora         35 | #########################                20.71%
Priateľská atmosféra    33 | ########################                 19.53%
```

---

# 9. Korelácia: lokalita × dôvod neúčasti

Popis: Napríklad zistíš **či ľudia zo Sásovej viac hovoria že je centrum ďaleko**.

```sql
SELECT
    loc.option_label AS location,
    reason_resp.option_label AS reason,
    COUNT(*) AS responses
FROM response_options loc
JOIN response_options reason_resp
    ON loc.response_id = reason_resp.response_id
WHERE loc.question_code = 'location'
AND reason_resp.question_code = 'not_attending_reason'
GROUP BY location, reason
ORDER BY responses DESC;
```

```
┌─────────────────────────────────────────────────────────────────┬───────────────────────────┬───────────┐
│                            location                             │          reason           │ responses │
│                             varchar                             │          varchar          │   int64   │
├─────────────────────────────────────────────────────────────────┼───────────────────────────┼───────────┤
│ BB Sásová (Rudlová, Sásová)                                     │ Nenavštevujeme            │        11 │
│ BB Fončorda, Radvaň                                             │ Kvôli vzdialenosti        │         8 │
│ BB Sásová (Rudlová, Sásová)                                     │ Nevyhovuje mi čas aktivít │         6 │
│ BB Stred (Banská Bystrica-centrum)                              │ Kvôli vzdialenosti        │         4 │
│ BB Stred (Banská Bystrica-centrum)                              │ Nevyhovuje mi čas aktivít │         3 │
│ BB Sásová (Rudlová, Sásová)                                     │ Nepáčia sa mi priestory   │         2 │
│ BB Sásová (Rudlová, Sásová)                                     │ Neatraktívne aktivity     │         2 │
│ BB Fončorda, Radvaň                                             │ Nenavštevujeme            │         1 │
│ BB Juh (lliaš, Kráľová, Kremnička, Pršianska Terasa, Rakytovce) │ Nevyhovuje mi čas aktivít │         1 │
│ BB Juh (lliaš, Kráľová, Kremnička, Pršianska Terasa, Rakytovce) │ Kvôli vzdialenosti        │         1 │
│ BB Východ (Majer, Senica, Šalková, Uhlisko)                     │ Nevyhovuje mi čas aktivít │         1 │
│ BB Východ (Majer, Senica, Šalková, Uhlisko)                     │ Kvôli vzdialenosti        │         1 │
│ BB Stred (Banská Bystrica-centrum)                              │ Nenavštevujeme            │         1 │
│ BB Fončorda, Radvaň                                             │ Nevyhovuje mi čas aktivít │         1 │
│ BB Juh (lliaš, Kráľová, Kremnička, Pršianska Terasa, Rakytovce) │ Neatraktívne aktivity     │         1 │
│ BB Sásová (Rudlová, Sásová)                                     │ Kvôli vzdialenosti        │         1 │
├─────────────────────────────────────────────────────────────────┴───────────────────────────┴───────────┤
│ 16 rows                                                                                       3 columns │
└─────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Lokalita × dôvod neúčasti (top kombinácie)
Sásová    × Nenavštevujeme         11 | ########################################
Fončorda  × Kvôli vzdialenosti      8 | #############################
Sásová    × Nevyhovuje čas          6 | ######################
Stred     × Kvôli vzdialenosti      4 | ###############
Stred     × Nevyhovuje čas          3 | ###########
```

---

# 10. Korelácia: vek dieťaťa × preferovaný čas

Popis: Napríklad či rodičia malých detí preferujú dopoludnie.

```sql
SELECT
    age.option_label AS age_group,
    time.option_label AS preferred_time,
    COUNT(*) AS responses
FROM response_options age
JOIN response_options time
    ON age.response_id = time.response_id
WHERE age.question_code = 'age'
AND time.question_code = 'preffered_time'
GROUP BY age_group, preferred_time
ORDER BY responses DESC;
```

```
┌────────────────────┬───────────────────────────────┬───────────┐
│     age_group      │        preferred_time         │ responses │
│      varchar       │            varchar            │   int64   │
├────────────────────┼───────────────────────────────┼───────────┤
│ 1-3 rokov          │ dopoludnia (napr. 9:00–12:00) │        25 │
│ 3-6 rokov          │ podvečer (po 17:00)           │        20 │
│ 3-6 rokov          │ popoludní (napr. 15:00–17:00) │        19 │
│ 1-3 rokov          │ popoludní (napr. 15:00–17:00) │        13 │
│ 3-6 rokov          │ víkendy                       │        12 │
│ 3-6 rokov          │ dopoludnia (napr. 9:00–12:00) │        10 │
│ 6+ rokov           │ podvečer (po 17:00)           │        10 │
│ 1-3 rokov          │ podvečer (po 17:00)           │        10 │
│ 6+ rokov           │ víkendy                       │         9 │
│ 1-3 rokov          │ víkendy                       │         9 │
│ 6+ rokov           │ popoludní (napr. 15:00–17:00) │         7 │
│ 0-1 rok            │ dopoludnia (napr. 9:00–12:00) │         7 │
│ 6+ rokov           │ dopoludnia (napr. 9:00–12:00) │         6 │
│ ešte sa nenarodilo │ dopoludnia (napr. 9:00–12:00) │         5 │
│ 0-1 rok            │ popoludní (napr. 15:00–17:00) │         5 │
│ ešte sa nenarodilo │ popoludní (napr. 15:00–17:00) │         5 │
│ ešte sa nenarodilo │ podvečer (po 17:00)           │         4 │
│ 0-1 rok            │ podvečer (po 17:00)           │         2 │
│ ešte sa nenarodilo │ víkendy                       │         1 │
├────────────────────┴───────────────────────────────┴───────────┤
│ 19 rows                                              3 columns │
└────────────────────────────────────────────────────────────────┘

Vek × preferovaný čas (top kombinácie)
1-3 rokov  × dopoludnia    25 | ########################################
3-6 rokov  × podvečer      20 | ################################
3-6 rokov  × popoludní     19 | ##############################
1-3 rokov  × popoludní     13 | ####################
3-6 rokov  × víkendy       12 | ###################
```

---

# 11. Korelácia: lokalita × preferovaný čas

Popis: Či rôzne časti mesta preferujú iný čas aktivít.

```sql
SELECT
    loc.option_label AS location,
    time.option_label AS preferred_time,
    COUNT(*) AS responses
FROM response_options loc
JOIN response_options time
    ON loc.response_id = time.response_id
WHERE loc.question_code = 'location'
AND time.question_code = 'preffered_time'
GROUP BY location, preferred_time
ORDER BY location, responses DESC;
```

```
┌─────────────────────────────────────────────────────────────────┬───────────────────────────────┬───────────┐
│                            location                             │        preferred_time         │ responses │
│                             varchar                             │            varchar            │   int64   │
├─────────────────────────────────────────────────────────────────┼───────────────────────────────┼───────────┤
│ BB Fončorda, Radvaň                                             │ dopoludnia (napr. 9:00–12:00) │        12 │
│ BB Fončorda, Radvaň                                             │ popoludní (napr. 15:00–17:00) │         7 │
│ BB Fončorda, Radvaň                                             │ podvečer (po 17:00)           │         6 │
│ BB Fončorda, Radvaň                                             │ víkendy                       │         5 │
│ BB Juh (lliaš, Kráľová, Kremnička, Pršianska Terasa, Rakytovce) │ podvečer (po 17:00)           │         2 │
│ BB Juh (lliaš, Kráľová, Kremnička, Pršianska Terasa, Rakytovce) │ popoludní (napr. 15:00–17:00) │         2 │
│ BB Juh (lliaš, Kráľová, Kremnička, Pršianska Terasa, Rakytovce) │ dopoludnia (napr. 9:00–12:00) │         1 │
│ BB Juh (lliaš, Kráľová, Kremnička, Pršianska Terasa, Rakytovce) │ víkendy                       │         1 │
│ BB Stred (Banská Bystrica-centrum)                              │ dopoludnia (napr. 9:00–12:00) │         7 │
│ BB Stred (Banská Bystrica-centrum)                              │ podvečer (po 17:00)           │         4 │
│ BB Stred (Banská Bystrica-centrum)                              │ víkendy                       │         3 │
│ BB Stred (Banská Bystrica-centrum)                              │ popoludní (napr. 15:00–17:00) │         1 │
│ BB Sásová (Rudlová, Sásová)                                     │ popoludní (napr. 15:00–17:00) │        27 │
│ BB Sásová (Rudlová, Sásová)                                     │ podvečer (po 17:00)           │        25 │
│ BB Sásová (Rudlová, Sásová)                                     │ víkendy                       │        16 │
│ BB Sásová (Rudlová, Sásová)                                     │ dopoludnia (napr. 9:00–12:00) │        15 │
│ BB Východ (Majer, Senica, Šalková, Uhlisko)                     │ dopoludnia (napr. 9:00–12:00) │         1 │
│ BB Východ (Majer, Senica, Šalková, Uhlisko)                     │ podvečer (po 17:00)           │         1 │
├─────────────────────────────────────────────────────────────────┴───────────────────────────────┴───────────┤
│ 18 rows                                                                                           3 columns │
└─────────────────────────────────────────────────────────────────────────────────────────────────────────────┘

Lokalita × preferovaný čas (top kombinácie)
Sásová    × popoludní      27 | ########################################
Sásová    × podvečer       25 | #####################################
Sásová    × víkendy        16 | ########################
Sásová    × dopoludnia     15 | ######################
Fončorda  × dopoludnia     12 | ##################
```

---

# 12. Korelácia: aktivity rodičov × aktivity detí

Popis: Aké kombinácie rodič + dieťa sú najčastejšie.

```sql
SELECT
    parent.option_label AS parent_activity,
    child.option_label AS child_activity,
    COUNT(*) AS responses
FROM response_options parent
JOIN response_options child
    ON parent.response_id = child.response_id
WHERE parent.question_code = 'activities_for_parent'
AND child.question_code = 'activities_for_children'
GROUP BY parent_activity, child_activity
ORDER BY responses DESC;
```

```
┌──────────────────────────────────────────────────────────┬─────────────────────────────────────────┬───────────┐
│                     parent_activity                      │             child_activity              │ responses │
│                         varchar                          │                 varchar                 │   int64   │
├──────────────────────────────────────────────────────────┼─────────────────────────────────────────┼───────────┤
│ prednášky / workshopy (výchova, psychológia, zdravie…)   │ herničky / voľná hra                    │        29 │
│ prednášky / workshopy (výchova, psychológia, zdravie…)   │ pohybové aktivity (cvičenie, tanec)     │        29 │
│ cvičenie pre mamičky (jóga, fit po pôrode, pilates…)     │ pohybové aktivity (cvičenie, tanec)     │        28 │
│ tvorivé dielne                                           │ pohybové aktivity (cvičenie, tanec)     │        25 │
│ tvorivé dielne                                           │ herničky / voľná hra                    │        23 │
│ praktické kurzy (napr. varenie, prvá pomoc)              │ pohybové aktivity (cvičenie, tanec)     │        23 │
│ prednášky / workshopy (výchova, psychológia, zdravie…)   │ hudobné aktivity / rytmika              │        23 │
│ tvorivé dielne                                           │ výtvarné / kreatívne aktivity           │        22 │
│ tvorivé dielne                                           │ hudobné aktivity / rytmika              │        21 │
│ prednášky / workshopy (výchova, psychológia, zdravie…)   │ výtvarné / kreatívne aktivity           │        21 │
│ prednášky / workshopy (výchova, psychológia, zdravie…)   │ adaptačné programy (príprava na škôlku) │        20 │
│ praktické kurzy (napr. varenie, prvá pomoc)              │ hudobné aktivity / rytmika              │        19 │
│ tvorivé dielne                                           │ jazykové aktivity                       │        18 │
│ prednášky / workshopy (výchova, psychológia, zdravie…)   │ jazykové aktivity                       │        17 │
│ praktické kurzy (napr. varenie, prvá pomoc)              │ výtvarné / kreatívne aktivity           │        17 │
│ relaxačné aktivity (jóga, dychové cvičenia, mindfulness) │ pohybové aktivity (cvičenie, tanec)     │        16 │
│ cvičenie pre mamičky (jóga, fit po pôrode, pilates…)     │ herničky / voľná hra                    │        16 │
│ praktické kurzy (napr. varenie, prvá pomoc)              │ jazykové aktivity                       │        15 │
│ cvičenie pre mamičky (jóga, fit po pôrode, pilates…)     │ adaptačné programy (príprava na škôlku) │        15 │
│ cvičenie pre mamičky (jóga, fit po pôrode, pilates…)     │ hudobné aktivity / rytmika              │        15 │
│ tvorivé dielne                                           │ adaptačné programy (príprava na škôlku) │        14 │
│ praktické kurzy (napr. varenie, prvá pomoc)              │ herničky / voľná hra                    │        14 │
│ cvičenie pre mamičky (jóga, fit po pôrode, pilates…)     │ výtvarné / kreatívne aktivity           │        14 │
│ podporné skupiny (diskusie, zdieľanie skúseností)        │ pohybové aktivity (cvičenie, tanec)     │        14 │
│ praktické kurzy (napr. varenie, prvá pomoc)              │ adaptačné programy (príprava na škôlku) │        13 │
│ relaxačné aktivity (jóga, dychové cvičenia, mindfulness) │ herničky / voľná hra                    │        12 │
│ relaxačné aktivity (jóga, dychové cvičenia, mindfulness) │ hudobné aktivity / rytmika              │        12 │
│ podporné skupiny (diskusie, zdieľanie skúseností)        │ herničky / voľná hra                    │        12 │
│ cvičenie pre mamičky (jóga, fit po pôrode, pilates…)     │ jazykové aktivity                       │        11 │
│ podporné skupiny (diskusie, zdieľanie skúseností)        │ hudobné aktivity / rytmika              │        11 │
│ podporné skupiny (diskusie, zdieľanie skúseností)        │ výtvarné / kreatívne aktivity           │        10 │
│ relaxačné aktivity (jóga, dychové cvičenia, mindfulness) │ adaptačné programy (príprava na škôlku) │        10 │
│ podporné skupiny (diskusie, zdieľanie skúseností)        │ adaptačné programy (príprava na škôlku) │         9 │
│ relaxačné aktivity (jóga, dychové cvičenia, mindfulness) │ výtvarné / kreatívne aktivity           │         8 │
│ relaxačné aktivity (jóga, dychové cvičenia, mindfulness) │ jazykové aktivity                       │         8 │
│ podporné skupiny (diskusie, zdieľanie skúseností)        │ jazykové aktivity                       │         7 │
├──────────────────────────────────────────────────────────┴─────────────────────────────────────────┴───────────┤
│ 36 rows                                                                                              3 columns │
└────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘

Rodič × dieťa (top kombinácie)
Prednášky   × Herničky        29 | ########################################
Prednášky   × Pohyb           29 | ########################################
Cvičenie    × Pohyb           28 | ######################################
Tvorivé     × Pohyb           25 | ##################################
Tvorivé     × Herničky        23 | ###############################
```

---

# 13. Heatmapa preferencií aktivít

Popis: Dá sa priamo použiť ako heatmap graf.

```sql
SELECT
    question_code,
    option_label,
    COUNT(*) AS responses
FROM response_options
GROUP BY question_code, option_label
ORDER BY question_code, responses DESC;
```

```
┌──────────────────────────────┬─────────────────────────────────────────────────────────────────┬───────────┐
│        question_code         │                          option_label                           │ responses │
│           varchar            │                             varchar                             │   int64   │
├──────────────────────────────┼─────────────────────────────────────────────────────────────────┼───────────┤
│ activities_for_children      │ pohybové aktivity (cvičenie, tanec)                             │        54 │
│ activities_for_children      │ herničky / voľná hra                                            │        46 │
│ activities_for_children      │ hudobné aktivity / rytmika                                      │        38 │
│ activities_for_children      │ adaptačné programy (príprava na škôlku)                         │        34 │
│ activities_for_children      │ výtvarné / kreatívne aktivity                                   │        31 │
│ activities_for_children      │ jazykové aktivity                                               │        31 │
│ activities_for_parent        │ prednášky / workshopy (výchova, psychológia, zdravie…)          │        45 │
│ activities_for_parent        │ tvorivé dielne                                                  │        35 │
│ activities_for_parent        │ cvičenie pre mamičky (jóga, fit po pôrode, pilates…)            │        35 │
│ activities_for_parent        │ praktické kurzy (napr. varenie, prvá pomoc)                     │        33 │
│ activities_for_parent        │ relaxačné aktivity (jóga, dychové cvičenia, mindfulness)        │        22 │
│ activities_for_parent        │ podporné skupiny (diskusie, zdieľanie skúseností)               │        21 │
│ age                          │ 1-3 rokov                                                       │        37 │
│ age                          │ 3-6 rokov                                                       │        36 │
│ age                          │ 6+ rokov                                                        │        19 │
│ age                          │ 0-1 rok                                                         │        11 │
│ age                          │ ešte sa nenarodilo                                              │         9 │
│ location                     │ BB Sásová (Rudlová, Sásová)                                     │        49 │
│ location                     │ BB Fončorda, Radvaň                                             │        19 │
│ location                     │ BB Stred (Banská Bystrica-centrum)                              │        11 │
│ location                     │ BB Juh (lliaš, Kráľová, Kremnička, Pršianska Terasa, Rakytovce) │         4 │
│ location                     │ BB Východ (Majer, Senica, Šalková, Uhlisko)                     │         2 │
│ not_attending_reason         │ Kvôli vzdialenosti                                              │        15 │
│ not_attending_reason         │ Nenavštevujeme                                                  │        12 │
│ not_attending_reason         │ Nevyhovuje mi čas aktivít                                       │        11 │
│ not_attending_reason         │ Neatraktívne aktivity                                           │         3 │
│ not_attending_reason         │ Nepáčia sa mi priestory                                         │         2 │
│ preffered_activity_frequency │ 1× týždenne                                                     │        55 │
│ preffered_activity_frequency │ 2–3× týždenne                                                   │        25 │
│ preffered_activity_frequency │ občas                                                           │        17 │
│ preffered_activity_frequency │ len pri špeciálnych akciách                                     │         9 │
│ preffered_cons               │ čas                                                             │        56 │
│ preffered_cons               │ cena                                                            │        45 │
│ preffered_cons               │ kvalita lektora                                                 │        35 │
│ preffered_cons               │ priateľská atmosféra                                            │        33 │
│ preffered_time               │ podvečer (po 17:00)                                             │        37 │
│ preffered_time               │ popoludní (napr. 15:00–17:00)                                   │        36 │
│ preffered_time               │ dopoludnia (napr. 9:00–12:00)                                   │        36 │
│ preffered_time               │ víkendy                                                         │        24 │
├──────────────────────────────┴─────────────────────────────────────────────────────────────────┴───────────┤
│ 39 rows                                                                                          3 columns │
└────────────────────────────────────────────────────────────────────────────────────────────────────────────┘

Top odpovede podľa tém

activities_for_children
  pohybové aktivity        54 | ########################################
  herničky / voľná hra     46 | ##################################
  hudobné / rytmika        38 | ############################
  adaptačné programy       34 | #########################
  jazykové aktivity        31 | #######################
  výtvarné aktivity        31 | #######################

activities_for_parent
  prednášky / workshopy    45 | ########################################
  tvorivé dielne           35 | ###############################
  cvičenie pre mamičky     35 | ###############################
  praktické kurzy          33 | #############################
  relaxačné aktivity       22 | ###################
  podporné skupiny         21 | ##################

preffered_cons
  čas                      56 | ########################################
  cena                     45 | ################################
  kvalita lektora          35 | #########################
  priateľská atmosféra     33 | ########################
```

---

# 14. Respondenti ktorí chcú spoločné aktivity

Popis: Percento ľudí ktorí chcú rodič + dieťa aktivity.

```sql
SELECT
    want_common_activities,
    COUNT(*) AS responses,
    ROUND(100.0 * COUNT(*) / SUM(COUNT(*)) OVER (), 2) AS percentage
FROM responses
GROUP BY want_common_activities;
```

```
┌────────────────────────┬───────────┬────────────┐
│ want_common_activities │ responses │ percentage │
│        varchar         │   int64   │   double   │
├────────────────────────┼───────────┼────────────┤
│ NULL                   │         1 │        1.2 │
│ áno                    │        76 │      91.57 │
│ nie                    │         6 │       7.23 │
└────────────────────────┴───────────┴────────────┘

Záujem o spoločné aktivity
Áno      76 | ######################################## 91.57%
Nie       6 | ###                                      7.23%
NULL      1 | #                                        1.20%
```

---

# 15. Najčastejšie kombinácie odpovedí

Popis: Silné signály v dátach (čo sa často vyskytuje spolu).

```sql
SELECT
    a.option_label AS option_a,
    b.option_label AS option_b,
    COUNT(*) AS responses
FROM response_options a
JOIN response_options b
    ON a.response_id = b.response_id
   AND a.option_code < b.option_code
GROUP BY option_a, option_b
ORDER BY responses DESC
LIMIT 20;
```

```
┌──────────────────────────────────────────────────┬────────────────────────────────────────────────────────┬───────────┐
│                     option_a                     │                        option_b                        │ responses │
│                     varchar                      │                        varchar                         │   int64   │
├──────────────────────────────────────────────────┼────────────────────────────────────────────────────────┼───────────┤
│ 1× týždenne                                      │ čas                                                    │        41 │
│ pohybové aktivity (cvičenie, tanec)              │ čas                                                    │        39 │
│ pohybové aktivity (cvičenie, tanec)              │ 1× týždenne                                            │        35 │
│ BB Sásová (Rudlová, Sásová)                      │ pohybové aktivity (cvičenie, tanec)                    │        34 │
│ prednášky / workshopy (výchova, psychológia, z…  │ čas                                                    │        33 │
│ herničky / voľná hra                             │ pohybové aktivity (cvičenie, tanec)                    │        33 │
│ 1× týždenne                                      │ cena                                                   │        33 │
│ herničky / voľná hra                             │ 1× týždenne                                            │        32 │
│ pohybové aktivity (cvičenie, tanec)              │ hudobné aktivity / rytmika                             │        32 │
│ herničky / voľná hra                             │ čas                                                    │        31 │
│ BB Sásová (Rudlová, Sásová)                      │ čas                                                    │        31 │
│ BB Sásová (Rudlová, Sásová)                      │ 1× týždenne                                            │        30 │
│ cena                                             │ čas                                                    │        30 │
│ prednášky / workshopy (výchova, psychológia, z…  │ pohybové aktivity (cvičenie, tanec)                    │        29 │
│ prednášky / workshopy (výchova, psychológia, z…  │ 1× týždenne                                            │        29 │
│ 3-6 rokov                                        │ 1× týždenne                                            │        29 │
│ herničky / voľná hra                             │ prednášky / workshopy (výchova, psychológia, zdravie…) │        29 │
│ podvečer (po 17:00)                              │ čas                                                    │        29 │
│ 3-6 rokov                                        │ čas                                                    │        29 │
│ cvičenie pre mamičky (jóga, fit po pôrode, pil…  │ pohybové aktivity (cvičenie, tanec)                    │        28 │
├──────────────────────────────────────────────────┴────────────────────────────────────────────────────────┴───────────┤
│ 20 rows                                                                                                     3 columns │
└───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘

Najčastejšie kombinácie odpovedí
1× týždenne + čas                  41 | ########################################
Pohybové aktivity + čas            39 | ######################################
Pohybové aktivity + 1× týždenne    35 | ##################################
Sásová + pohybové aktivity         34 | #################################
Prednášky + čas                    33 | ################################
Herničky + pohybové aktivity       33 | ################################
1× týždenne + cena                 33 | ################################
```

---

## Hlavné závery

1. Kto odpovedal
   - Dominantná lokalita je Sásová
   - Najsilnejší vekový segment detí je 1–6 rokov

2. Čo chcú
   - Pre deti vedú pohybové aktivity, herničky, hudobné aktivity
   - Pre rodičov vedú prednášky/workshopy, potom cvičenie, tvorivé dielne, praktické kurzy

3. Kedy to chcú
   - Najsilnejší dopyt je po podvečeri, ale veľmi blízko sú aj popoludnia a dopoludnia
   - Segment 1–3 roky preferuje skôr dopoludnie
   - Segment 3–6 rokov skôr popoludnie / podvečer

4. Čo ich brzdí
   - Najväčšia bariéra je vzdialenosť
   - Druhá je že miesto nenavštevujú
   - Tretia je nevhodný čas aktivít

5. Čo rozhoduje pri výbere
   - čas
   - cena
   - kvalita lektora
   - priateľská atmosféra

6. Veľmi silný signál
   - 91.57 % respondentov chce spoločné aktivity rodič + dieťa

## Odporúčanie na pilotný program

Na základe dát by dávalo zmysel otestovať:

- 1× týždenne pravidelný blok
  - rodičia: prednáška / workshop alebo praktický kurz
  - deti: pohybová aktivita alebo hernička
- rozdeliť termíny minimálne na dva segmenty
  1.) dopoludnie pre menšie deti
  2.) popoludnie / podvečer pre vek 3–6 rokov
- komunikačne zdôrazniť
  1.) dobrý čas
  2.) prijateľnú cenu
  3.) praktický prínos aktivít
