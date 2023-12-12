export interface ResultSet {
	items: Array<ResultItem>
}

export interface ResultItem {
	sim: number
	sig: string
	summary: string
	link: string
}

const host = import.meta.env.MODE === 'production'
	? import.meta.env.VITE_EXTERN_ENDPOINT
	: '/api'

console.log('host', host)
console.log('mode', import.meta.env)

export const query = async (
	query: string,
): Promise<ResultSet> => {
	const raw = await fetch(`${host}/search`, {
		method: 'POST',
		body: JSON.stringify({ q: query }),
	})

	return await raw.json()
}
