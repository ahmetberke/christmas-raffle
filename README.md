# What is this?
This application is a tool that allows you to prepare mail designs specific to your group and make Christmas raffles.

# How is it working?
It performs a draw by randomly matching between the participants you add to the config file (people cannot go out to themselves). And And result is sent to that person as an e-mail.

# How do I/We Use?
- step 1: You need an email address to send mail (ex: mychristmasraffle.gmail.com)
- step 2: Your e-mail account must be open for sending e-mails by applications. For Gmail: [click here](https://kb.synology.com/en-global/SRM/tutorial/How_to_use_Gmail_SMTP_server_to_send_emails_for_SRM)
- step 3: Design your own custom christmas card or use the default design. (christmas card must be a single html file) `default: ./web/template/index.html` 
- step 4: Open the configs yaml file and adjust the settings. (sample settings will appear)

## How to Run?
run on terminal
```
go install
```
``` 
go run ./cmd/raffle/main.go
```
