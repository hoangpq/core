---
title: Receive email when there is an ERC20 transfer
description: >-
  Tutorial: How to create an application that sends an email every time there is
  a transfer on an ERC20 smart contract.
published_link: 'https://docs.mesg.com/tutorials/erc20-transfer-notifications/receive-email-when-there-is-an-erc20-transfer.html'
---

# Receive email when there is an ERC20 transfer

Today we'll learn how to create a JavaScript application that connects the two previously created MESG Services:

* [Tutorial: Transfer notifications from an ERC20 transfer](./listen-for-transfers-of-an-ethereum-erc20-token.md)
* [Tutorial: Sending emails through SendGrid](./send-emails-with-sendgrid.md)

You can access the final version of the [source code on GitHub](https://github.com/mesg-foundation/core/tree/master/docs/tutorials/erc20-transfer-notifications/email-notification-on-erc20-transfer).

::: tip
MESG Core should already be installed on your computer. If it isn't yet, [install it here](../../guide/start-here/installation.html).
:::

## Setup

### Start core

Let's start MESG Core, if it isn't already running:

```text
mesg-core start
```

### Deploy the Services

We'll start with our two already-created services. If you haven't already, make sure to read the tutorials to see how they were created.

#### Deploy the ERC20 Service

```bash
mesg-core service deploy ./PATH_TO_THE_ERC20_SERVICE
```

Make sure to copy/paste the service ID somewhere. You will need it later.

#### Deploy the SendGrid Service

This will be the same process as previous service:

```bash
mesg-core service deploy ./PATH_TO_THE_SENDGRID_SERVICE
```

Make sure to copy/paste the service ID somewhere. You will need it later.

### Get a SendGrid API Key

Go to [https://app.sendgrid.com/settings/api\_keys](https://app.sendgrid.com/settings/api_keys) then click on "Create API Key". Select "Full Access" and follow the steps. Copy/paste the generated API Key somewhere. You will need it later.

### Init the application

Let's init the application:

```bash
npm init -y
```

Then, install `mesg-js` library:

```bash
npm install --save mesg-js
```

## Code the application

The setup for our application has finished. Now, let's code it. Create and open a `index.js` file.

### Listen for Transfer events

Let's define a variable for the event we want to listen to:

```javascript
// Event we need to listen
const erc20Transfer = {
  serviceID: '__ERC20_SERVICE_ID__', // The serviceID of the ERC20 service deployed
  eventKey: 'transfer' // The event we want to listen
}
```

Replace `__ERC20_SERVICE_ID__` with service ID of the ERC20 service. You can deploy the ERC20 service again if you didn't copy its service ID.

### Execute the send email task

Let's define another variable for the task we want to execute:

```javascript
// Task to execute
const sendEmail = {
  serviceID: '__SENDGRID_SERVICE_ID__', // The serviceID of the service to send emails
  taskKey: 'send', // The task we want to execute
  inputs: (eventKey, eventData) => { // This function returns the inputs for of task send based on the data of the event
    console.log('New ERC20 transfer received. will send an email. Transaction hash:', eventData.transactionHash)
    return {
      apiKey: '__SENDGRID_API_KEY__',
      from: 'test@erc20notification.com',
      to: '__REPLACE_WITH_YOUR_EMAIL__',
      subject: 'New ERC20 transfer',
      text: `Transfer from ${eventData.from} to ${eventData.to} of ${eventData.value} tokens -> ${eventData.transactionHash}`
    }
  }
}
```

Replace `__SENDGRID_SERVICE_ID__`, `__SENDGRID_API_KEY__` and `__REPLACE_WITH_YOUR_EMAIL__` by the right values. You can deploy the SendGrid service again if you didn't copy its service ID.

#### Link the event to the task

Great, now we need to link this event and task together using MESG.

Require the `mesg-js` library on top of your code with the following code:

```javascript
const MESG = require('mesg-js').application()
```

Then simply add the `whenEvent` function at the end of your code to link the event to the task:

```javascript
MESG.whenEvent(erc20Transfer, sendEmail)
console.log('Listening ERC20 transfer...')
```

#### Run it!

Now your application is ready.

```bash
node index.js
```

Your application will automatically start the services, connect to the Ethereum network, and send you an email every time a transfer occurs on any ERC20 token.

As it is based on the ERC20 transfer activity, it could take a while to receive the first email. You can check the logs of the service ERC20 by running the command (replace `__ERC20_SERVICE_ID__` with service ID of the ERC20 service):
```
mesg-core service logs __ERC20_SERVICE_ID__
```

Be careful, ERC20 tokens have a lot of activity so it is possible to have thousands of emails per day and reach the SendGrid limit if you leave your application running.


### Final version of the source code

<card-link url="https://github.com/mesg-foundation/core/tree/master/docs/tutorials/erc20-transfer-notifications/email-notification-on-erc20-transfer"></card-link>

