// let copyText = document.querySelector('.copy-text')
// copyText.querySelector('button').addEventListener('click', function () {
// 	let input = copyText.querySelector('input.text')
// 	input.select()
// 	document.execCommand('copy')
// 	copyText.classList.add('active')
// 	window.getSelection().removeAllRanges()
// 	setTimeout(function () {
// 		copyText.classList.remove('active')
// 	}, 2500)
// })

const form = document.getElementById('shorten-form')


form.addEventListener('submit', (event) => {
	// handle the form data
	event.preventDefault()
	const data = event.target
	console.log(event)
})
