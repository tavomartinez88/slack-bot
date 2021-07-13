# slack-bot

![](https://img.shields.io/static/v1?label=Aws%20&message=3.27&color=informational)
![](https://img.shields.io/github/go-mod/go-version/tavomartinez88/slack-bot/main)
![](https://img.shields.io/static/v1?label=Terraform%20&message=0.13.2&color=informational&)
![](https://img.shields.io/static/v1?label=gomock%20&message=1.6.0&color=informational)
![](https://img.shields.io/github/languages/count/tavomartinez88/slack-bot)
![](https://img.shields.io/github/languages/top/tavomartinez88/slack-bot)
![](https://img.shields.io/github/issues-pr/tavomartinez88/slack-bot)
![](https://img.shields.io/github/v/tag/tavomartinez88/slack-bot?include_prereleases)
![](https://img.shields.io/github/contributors/tavomartinez88/slack-bot)
![](https://img.shields.io/github/last-commit/tavomartinez88/slack-bot)
### Tech used:
- Go >= 1.15
- AWS(Dynamo-Lambda-Sns-Chatbot)
- Slack
- Terraform >= v0.13.2
- gomock >= v1.6.0

### Build Lambda
cd lambda/slack-bot-release-aws-lambda/ && sh build.sh

### Terraform
- cd terraform
    - set profile and region in main.tf
    - terraform init
    - terraform plan
    - terraform apply --auto-approve=true   

### Create Chatbot    
-  You should access to aws chatbot under aws console
-  Choose Slack as client
    -  Choose the workspace in dropdown list
-  Configure new channel
    - On details set any name, e.g. bot-para-notificar-releases
        - Note: not check logging. it apply charge
    - Under Channel type, choose Private
    - Enter channel ID (e.g. T026YK7T0Q5)
        - if you haven't any channel, you shoud go to slack for create channel
        - if you have channel for it, enter id. (On my case it's C0271L15MBL)
    - Permissions
        -  Set name e.g. bot-para-notificar-releases-role
        -  Policy templates :
            -  Read-only command permissions
            -  Lambda-invoke command permissions
    - Notifications
        - select region sns topics(on my case, it's us-east-1)  
        - select topic  
-   Finally, click on configure button

### How do you should use it?
@aws lambda invoke --function-name [name] --region [region] --payload [json]

for example:
```
@aws lambda invoke --function-name slack-release-service-dev-da23 
--region us-east-1 
--payload {
    “title”: “titulo”,
    “description”: “se realizo refacto de ...”,
    “product”: “MyProduct”,
    “detail”: “https://any-endpoint.com”,
    “team”: “MyTeam”,
    “status”: “IMPLEMENTADO”,
    “owner”: “Gustavo Martinez”,
    “result”: “EXITOSO”,
    “observations”: “sin observaciones”
}  
```         

### Documentation
![alt text](docs/Slack-bot.png?raw=true)   