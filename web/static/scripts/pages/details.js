import {form} from "../utils/element.js"
import {doReq} from "../components/req.js"

form.addEventListener('submit', async evt => evt.preventDefault())

form.querySelector('button[type=submit]').addEventListener('click', async evt => {
    const res = await doReq('PUT', {
        name: form.querySelector('input[type=text]').value,
        finalGrade: Number(form.querySelector('#grade').value),
        schoolCohortOf: Number(form.querySelector('#cohort').value)
    })
    if (res.status === 200) {
        alert('Student updated successfully!')
        location.reload()
    } else {
        const { message } = await res.json()
        confirm(message + '. Do you want to reload browser?') && location.reload()
    }
})

document.getElementById('del-btn').addEventListener('click', async evt => {
    if (confirm('Are you sure you want to delete this student?')) {
        const res = await doReq('DELETE', null)
        if (res.status === 204) {
            alert('Student deleted successfully!')
            location.pathname = '/'
        } else {
            const { message } = await res.json()
            confirm(message + '. Do you want to reload browser?') && location.reload()
        }
    }
})