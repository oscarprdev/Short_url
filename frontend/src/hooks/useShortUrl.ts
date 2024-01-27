import { useMutation } from '@tanstack/react-query';
import { ShortUrlInput, shortUrl } from '../services/api/shortUrl';
import { useGlobalStore } from '../store/globalState';

export const useShortUrl = () => {
	const { addUrl } = useGlobalStore();

	return useMutation({
		mutationFn: async ({ originalUrl, userId }: ShortUrlInput) => {
			return await shortUrl({ originalUrl, userId });
		},
		onSuccess: ({ url }) => {
			addUrl(url);
		},
	});
};
