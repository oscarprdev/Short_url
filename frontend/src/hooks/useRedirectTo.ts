import { useUrlStore } from '../store/urlStore/urlStore';

export const useRedirectTo = (id?: string) => {
	const { urls } = useUrlStore();

	return (id && urls && urls.find((url) => url.titleUrl === id)?.originalUrl) || '/';
};
