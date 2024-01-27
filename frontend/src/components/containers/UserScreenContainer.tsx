import ShortUrlForm from '../ShortUrlForm';
import UrlCardList from '../UrlCardList';
import { useShortUrl } from '../../hooks/useShortUrl';
import { useGlobalStore } from '../../store/globalState';

const UserScreenContainer = () => {
	const { user, urls } = useGlobalStore();
	const { mutate: addUrl } = useShortUrl();

	return (
		<main className='flex flex-col h-full w-full items-center gap-20 z-10 mt-28'>
			<ShortUrlForm
				addUrl={addUrl}
				userId={user?.id}
			/>
			{urls && <UrlCardList urls={urls} />}
		</main>
	);
};

export default UserScreenContainer;
