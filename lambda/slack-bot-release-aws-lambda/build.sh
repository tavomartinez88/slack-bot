rm -rf ./../bin &&\
mkdir ./../bin &&\
GOOS=linux go build -o ./../bin/slack-bot-release-aws-lambda &&\
cd ./../bin &&\
zip slack-bot-release.zip slack-bot-release-aws-lambda