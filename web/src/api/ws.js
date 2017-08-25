import config from '@/config'

const ws = new WebSocket(
    `${config.service.ws}`
)

ws.onopen = () => console.log('WS open')

// Register 'messageKey' => callback()
const messageRegister = {};

ws.onmessage = ({ data }) => {
    if (data.length === 0) return
    try {
        const message = JSON.parse(data)

        if (messageRegister[message.type]) {
            console.info(`Dispatching ${message.type}`)
            messageRegister[message.type].forEach(cb => cb())
        }
    } catch (e) {
        console.error(e)
    }
}

ws.registerFor = (message, callback) => {
    if (!messageRegister[message]) {
        messageRegister[message] = []
    }

    return messageRegister[message].push(callback)
}

export default ws
