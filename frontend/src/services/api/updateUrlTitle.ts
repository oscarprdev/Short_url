import { API_URL } from '../../constants/apiUrl';
import { Url } from '../../types/url';

export interface UpdateUrlTitleInput {
	userId?: string;
	urlId: string;
	urlTitle: string;
}

export interface UpdateUrlTittleOutput {
	status: number;
	url: Url;
}

export const updateUrlTitle = async ({ userId, urlId, urlTitle }: UpdateUrlTitleInput): Promise<UpdateUrlTittleOutput> => {
	const response = await fetch(`${API_URL}/url/title?id=${userId}`, {
		method: 'POST',
		body: JSON.stringify({
			urlId,
			urlTitle,
		}),
	});

	const jsonData = await response.json();

	if (jsonData.status !== 200) {
		throw new Error(jsonData.details);
	}

	return jsonData;
};
