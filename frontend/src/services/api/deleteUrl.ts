import { API_URL } from '../../constants/apiUrl';

export interface DeleteUrlInput {
	userId?: string;
	urlId: string;
}

export interface DeleteUrlTittleOutput {
	status: number;
	message: string;
}

export const deleteUrl = async ({ userId, urlId }: DeleteUrlInput): Promise<DeleteUrlTittleOutput> => {
	const response = await fetch(`${API_URL}/url?id=${userId}&urlId=${urlId}`, {
		method: 'DELETE',
	});

	const jsonData = await response.json();

	if (jsonData.status !== 200) {
		throw new Error(jsonData.details);
	}

	return jsonData;
};
