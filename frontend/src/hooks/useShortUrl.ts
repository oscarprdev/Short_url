import { useMutation } from '@tanstack/react-query';
import { ShortUrlInput, shortUrl } from '../services/api/shortUrl';
import { useGlobalStore } from '../store/globalState';

export const useShortUrl = () => {
	const { addUrl, setError, user, clearStore } = useGlobalStore();

	return useMutation({
		mutationFn: async ({ originalUrl, userId }: ShortUrlInput) => {
			return await shortUrl({ originalUrl, userId });
		},
		onSuccess: ({ url }) => {
			if (user) {
				addUrl(url);
			} else {
				clearStore();
				addUrl(url);
			}
		},
		onError: (error) => {
			setError(error.message);
		},
	});
};
