const MESG = require('mesg-js').service()
const axios = require('axios')
const bodyParser = require('body-parser')
const app = require('express')()
const cors = require('cors')
app.use(cors())
app.use(bodyParser.urlencoded({ extended: false }))
app.use(bodyParser.json())

const call = ({ url, data, headers }, { error, result }) => axios({
  method: 'POST',
  url,
  headers,
  data
})
  .then(x => result(x.data))
  .catch(e => error({ message: e.toString() }))

MESG.listenTask({ call })

app.post('/webhook', async (req, res) => {
  const response = await MESG.emitEvent('request', {
    headers: req.headers,
    data: req.body,
  })
  res.send(response)
})

app.listen(3000, () => console.log('Example app listening on port 3000!'))