import { useEffect } from 'react';
import { useUrlUsage } from '../hooks/useUrlUsage';
import { useGlobalStore } from '../store/globalState';
import { useModal } from '../hooks/useModal';
import RedirectingModal from '../components/modals/RedirectingModal';

interface RedirectScreenProps {
	shortUrl: string;
}

const RedirectScreen = ({ shortUrl }: RedirectScreenProps) => {
	const { urls } = useGlobalStore();
	const { mutate: urlUsage } = useUrlUsage();
	const modal = useModal();

	useEffect(() => {
		const originalUrl = urls?.find((url) => url.shortUrl.match(shortUrl));

		if (originalUrl) {
			urlUsage({ urlId: originalUrl?.id });

			modal.openModal(<RedirectingModal />);

			window.location.href = originalUrl.originalUrl || '/';
		}
		// eslint-disable-next-line react-hooks/exhaustive-deps
	}, [shortUrl]);

	return <></>;
};

export default RedirectScreen;
