import { API_URL } from '../../constants/apiUrl';
import { Url } from '../../types/url';

export interface ShortUrlInput {
	originalUrl: string;
	userId?: string;
}

export interface ShortUrlOutput {
	status: number;
	url: Url;
}

export const shortUrl = async ({ originalUrl, userId }: ShortUrlInput): Promise<ShortUrlOutput> => {
	const url = userId ? `${API_URL}/url?id=${userId}` : `${API_URL}/url`;

	const data = await fetch(url, {
		method: 'POST',
		body: JSON.stringify({ originalUrl }),
	});

	const jsonData = await data.json();

	if (jsonData.status !== 200) {
		throw new Error(jsonData.details);
	}

	return jsonData;
};
