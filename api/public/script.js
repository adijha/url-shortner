function validateUrl(url) {
	var regexp =
		/(ftp|http|https):\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?/
	return regexp.test(url)
}

function makeRequest(requestOptions) {
	fetch('http://localhost/api/v1/', requestOptions)
		.then((response) => response.text())
		.then((result) => console.log(result))
		.catch((error) => console.log('error', error))
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

	console.log(payload)
	const response = await makeRequest(payload)

	console.log(response)

	// if (response.status === 200) {
	// 	const data = await response.json()
	// 	const shortUrl = data.shortUrl
	// 	const shortUrlEl = document.getElementById('short-url')
	// 	shortUrlEl.innerHTML = shortUrl
	// 	shortUrlEl.href = shortUrl
	// 	shortUrlEl.style.display = 'block'
	// }
})
