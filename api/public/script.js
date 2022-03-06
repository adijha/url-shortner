function validateUrl(url) {
	var regexp =
		/(ftp|http|https):\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?/
	return regexp.test(url)
}

function makeRequest(method, url, callback) {
	var xhr = new XMLHttpRequest()
	xhr.open(method, url, true)
	xhr.onreadystatechange = function () {
		if (xhr.readyState == 4) {
			callback(xhr.responseText)
		}
	}
	xhr.send()
}

// for (let i = 0; i < 3; i++) {
// 	const element = form.elements[i]
// 	console.log(element)
// }

function generatePayload(elements) {
	const body = {}
	for (let i = 0; i < 3; i++) {
		body[elements[i].name] = elements[i].value
	}
	if (!validateUrl(body['long-url'])) {
		alert('Please enter  a valid URL')
		return
	}

	var payload = {
		url: '/api/v1',
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body,
	}
	return payload
}

const form = document.getElementById('shorten-form')

form.addEventListener('submit', async (event) => {
	// handle the form data
	event.preventDefault()

	// const longUrl = form.elements['long-url'].value
	// const slug = form.elements['slug'].value
	// const expiry = form.elements['expiry'].value

	// check if the url is valid
	// if (!validateUrl(longUrl)) {
	// 	alert('Please enter a valid URL')
	// 	return
	// }

	const payload = generatePayload(form.elements)

	console.log(payload)
	return
	const response = await makeRequest('POST', '/api/v1/shorten', payload)

	console.log(response)

	if (response.status === 200) {
		const data = await response.json()
		const shortUrl = data.shortUrl
		const shortUrlEl = document.getElementById('short-url')
		shortUrlEl.innerHTML = shortUrl
		shortUrlEl.href = shortUrl
		shortUrlEl.style.display = 'block'
	}
})
