function validateUrl(url) {
	var regexp =
		/(ftp|http|https):\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?/
	return regexp.test(url)
}

const makeRequest = async (requestOptions) =>
	await fetch('http://localhost/api/v1/', requestOptions)

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
	try {
		const response = await makeRequest(payload)
		const result = await response.json()
		const shortUrl = result?.short
		if (shortUrl) {
			const shortUrlEl = document.getElementById('short-url')
			const shortUrlWrapper = document.getElementById('short-url-wrapper')
			shortUrlEl.innerHTML = shortUrl
			shortUrlEl.value = shortUrl
			shortUrlEl.href = shortUrl
			shortUrlWrapper.style.display = 'block'
		}
	} catch (error) {
		alert(error)
	}
})
