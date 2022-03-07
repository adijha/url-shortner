function validateUrl(url) {
	var regexp =
		/(ftp|http|https):\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?/
	return regexp.test(url)
}

function showShortURL(result) {
	if (result.short) {
		var shortURL = document.getElementById('short-url')
		shortURL.innerHTML = result.short
	} else {
		alert('Error, please try again')
	}
}

async function makeRequest(requestOptions) {
	// getDevices = async () => {
	// 	const location = window.location.hostname
	// 	const settings = {
	// 		method: 'POST',
	// 		headers: {
	// 			Accept: 'application/json',
	// 			'Content-Type': 'application/json',
	// 		},
	// 	}
	// 	try {
	// 		const fetchResponse = await fetch(
	// 			`http://${location}:9000/api/sensors/`,
	// 			settings
	// 		)
	// 		const data = await fetchResponse.json()
	// 		return data
	// 	} catch (e) {
	// 		return e
	// 	}
	// }

	// fetch('http://localhost/api/v1/', requestOptions)
	// 	.then((response) => response.text())
	// 	.then((result) => result)
	// 	.catch((error) => console.log('error', error))

	return await fetch('http://localhost/api/v1/', requestOptions)
}

function generatePayload(elements) {
	const body = {}
	for (let i = 0; i < 3; i++) {
		body[elements[i].name] = elements[i].value
	}
	if (!validateUrl(body['url'])) {
		alert('Please enter  a valid URL')
		return
	}

	var myHeaders = new Headers()
	myHeaders.append('Content-Type', 'application/json')

	// var raw = JSON.stringify(body)
	console.log(elements[2].value)
	var raw = JSON.stringify({
		url: elements[0].value,
		short: elements[1].value,
		expiry: Number(elements[2].value),
	})
	var payload = {
		method: 'POST',
		headers: myHeaders,
		body: raw,
		redirect: 'follow',
	}
	return payload
}

const form = document.getElementById('shorten-form')

form.addEventListener('submit', async (event) => {
	// handle the form data
	event.preventDefault()

	const payload = generatePayload(form.elements)

	const response = await makeRequest(payload)
	const result = await response.json()

	const shortUrl = result?.short
	console.log(shortUrl)
	if (shortUrl) {
		const shortUrlEl = document.getElementById('short-url')
		const shortUrlWrapper = document.getElementById('short-url-wrapper')
		shortUrlEl.innerHTML = shortUrl
		shortUrlEl.value = shortUrl
		shortUrlEl.href = shortUrl
		shortUrlWrapper.style.display = 'block'
	}
})
