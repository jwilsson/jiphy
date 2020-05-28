# Jiphy
A handy Slack bot to send your favorite GIFs.

## Prerequisites
* A Slack workspace.
* An AWS account.

## Setup
1. Start by creating a [Slack app and setting up a Slash Command](https://api.slack.com/interactivity/slash-commands).
2. Configure your [AWS account to work with Serverless](https://serverless.com/framework/docs/providers/aws/guide/credentials/).
3. Configure a `DYNAMO_TABLE_NAME` environmental variable with the name of your DynamoDB table.
4. Configure a `SLACK_SIGNING_SECRET` environmental variable with your Slack app's signing secret.
5. After deploying the first time, add some images to your DynamoDB table, see the item structure below.
6. Change any other values in `serverless.yml` to fit your needs.
7. Profit!

### Example image item
```json
{
    "giphy_url": "https://giphy.com/gifs/studiosoriginals-gloria-domitille-collardey-business-woman-3o7TKy1qgGdbbMalcQ",
    "image_url": "https://media.giphy.com/media/3o7TKy1qgGdbbMalcQ/giphy.gif",
    "image_name": "tubular"
}
```

* `giphy_url` - points to a Giphy page.
* `image_url` - points to the actual image file, the "social" size is recommended for Slack to always show it.
* `image_name` - name of the image. Used in the `/jiphy` command.

## Usage
To send one of the available GIFs, type `/jiphy <image>` where `<image>` points to one of the `image_name`s in your DynamoDB table. There is also a `/jiphy list` command to list all available images, only shown to the one sending it.
