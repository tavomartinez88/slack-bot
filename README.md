# slack-bot

### Tech used:
- Go >= 1.15
- AWS(Dynamo-Lambda-Sns-Chatbot)
- Slack
- Terraform >= v0.13.2

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
    - Enter channel ID
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
    “description”: “se corrigio algo”,
    “product”: “Debit”,
    “detail”: “https://any-endpoint.com”,
    “team”: “Transaccional”,
    “status”: “IMPLEMENTADO”,
    “owner”: “Gustavo Martinez”,
    “result”: “EXITOSO”,
    “observations”: “sin observaciones”
}  
```         

### Documentation
![alt text](docs/Slack-bot.png?raw=true)

### How continue?
- Integrate Terragrunt
- Left automatize creation chatbot using terraform, today's manually
- Add Unit Test
- Add more logs         