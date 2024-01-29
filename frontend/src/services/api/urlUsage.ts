import { API_URL } from '../../constants/apiUrl';
import { Url } from '../../types/url';

export interface UrlUsageInput {
	urlId: string;
	userId?: string;
}

export interface UrlUsageOutput {
	status: number;
	url: Url;
}

export const urlUsage = async ({ userId, urlId }: UrlUsageInput): Promise<UrlUsageOutput> => {
	const response = await fetch(`${API_URL}/url/use?id=${userId}`, {
		method: 'POST',
		body: JSON.stringify({
			urlId,
		}),
	});

	const jsonData = await response.json();

	if (jsonData.status !== 200) {
		throw new Error(jsonData.details);
	}

	return jsonData;
};
