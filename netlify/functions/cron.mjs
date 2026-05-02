export default async (_) => {
    await fetch(process.env.URL + '/.netlify/functions/job', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'X-Api-Key': process.env.API_KEY,
        },
        body: JSON.stringify({}),
    })
}
