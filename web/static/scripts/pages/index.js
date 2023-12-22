import {doReq} from "../components/req.js"
import {form} from "../utils/element.js"

form.addEventListener('submit', async evt => {
    evt.preventDefault()

    const res = await doReq('POST', {
        name: form.querySelector('input[type=text]').value,
        finalGrade: Number(form.querySelector('#grade').value),
        schoolCohortOf: Number(form.querySelector('#cohort').value)
    })

    if (res.status === 201) {
        alert('Student created successfully!')
        location.reload()
    } else {
        const { message } = await res.json()
        confirm(message + '. Do you want to reload browser?') && location.reload()
    }
})


document.getElementById('del-all-btn') && document.getElementById('del-all-btn').addEventListener('click', async evt => {
    if (confirm('Are you sure you want to delete all students?')) {
        const res = await doReq('DELETE', null)
        if (res.status === 204) {
            alert('All students deleted successfully!')
            location.reload()
        } else {
            const { message } = await res.json()
            confirm(message + '. Do you want to reload browser?') && location.reload()
        }
    }
})