import { useMutation } from '@tanstack/react-query';
import { shortUrl } from '../services/api/shortUrl';
import { useUrlStore } from '../store/urlStore/urlStore';

export const useShortUrl = () => {
	const { setUrl } = useUrlStore();

	return useMutation({
		mutationFn: async (url: string) => {
			return await shortUrl({ originalUrl: url });
		},
		onSuccess: ({ url }) => {
			setUrl(url);
		},
	});
};
