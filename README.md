# Mcmamina

## Development

You have different options for the local development. You can install dependencies and run locally.
  
  - run in docker using `docker compose up --watch`
  - run it directly on local machine (you need to have postgresql, run migration manually,...)

## Deployment



## Env variables

### GOOGLE_API_KEY

To create API KEY:

    - go to google console > Enabled API & service
    - Select [Credentials](https://console.cloud.google.com/apis/credentials?project=mcmamina&pli=1)
    - Click on "Create Credential" button. Select API KEY.
    - Copy API KEY and provide it to `.env``

### GOOGLE_CALENDAR_ID
This should be taken from the published mcmamina calendar.

    - sign in to google.
    - go to google calendar.
    - select calendar you want to use and go to it's settings
    - locate "Identifikátor kalendára" or calendar identifier
    - Copy and 
    
### GOOGLE_SMTP_PWD

Note: this one is not used in the app.

    - go to https://myaccount.google.com/
    - search for "App passwords"
    - create new password
    - copy and use

### GOOGLE_SMTP_MAIL

Assign the email to be used (mcmamina@mcmamina.sk).

### SESSION_KEY
To generate random session key, you can use this command:

```sh
LC_ALL=C tr -dc 'A-Za-z0-9' </dev/urandom | head -c 12 ; echo
```

### GOOGLE_CAPTCHA_SITE

    - go to https://console.cloud.google.com/ 
    - click on "View All Products" (or menu button on upper left corner next to Google Cloud)
    - Select "Security" (you need to go into the menu)
    - Look for "Detection and Controls" > then click reCAPTCHA
    - Create Key (Follow the instructions, platform type is Website)
    - Copy the ID and assign it to GOOGLE_CAPTCHA_SITE

### GOOGLE_AUTH_CLIENT_ID, GOOGLE_AUTH_CLIENT_SECRET and GOOGLE_AUTH_REDIRECT_PATH

    - go to https://console.cloud.google.com/ 
    - click on "View All Products" (or menu button on upper left corner next to Google Cloud)
    - Select "Google Auth Platform"
    - Create new client if not exists
    - Copy "Client ID" and "Client secret"


### POSTGRES_DB, POSTGRES_USER and POSTGRES_PASSWORD

You need to setup the database.

### PANORAMA_URL

The default Panorama url is "http://mcmamina.panfoto.sk/"