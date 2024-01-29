import { API_URL } from '../../constants/apiUrl';
import { Url } from '../../types/url';
import { User } from '../../types/user';

export interface GetUserInfoInput {
	userId: string;
}

export interface GetUserInfoOutput {
	status: number;
	user: User;
	urls: Url[];
}

export const getUserInfo = async ({ userId }: GetUserInfoInput): Promise<GetUserInfoOutput> => {
	const data = await fetch(`${API_URL}/users/describe?id=${userId}`);

	const jsonData = await data.json();

	console.log(jsonData);

	return jsonData;
};
