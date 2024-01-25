import { API_URL } from '../../constants/apiUrl';
import { Url } from '../../types/url';

export interface ShortUrlInput {
	originalUrl: string;
}

export interface ShortUrlOutput {
	status: number;
	url: Url;
}

export const shortUrl = async ({ originalUrl }: ShortUrlInput): Promise<ShortUrlOutput> => {
	const data = await fetch(`${API_URL}/url`, {
		method: 'POST',
		body: JSON.stringify({ originalUrl }),
	});

	const jsonData = await data.json();

	return jsonData;
};
