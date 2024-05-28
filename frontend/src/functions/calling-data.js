let getData = (url) => {
  let apiUrl = 'http://localhost:8000/api/' + url
  fetch(apiUrl)
  .then((response) => {
    return response.json()
  })
  .then((data) => {
    return data.data
  })
}