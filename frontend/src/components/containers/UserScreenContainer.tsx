import ShortUrlForm from '../ShortUrlForm';
import UrlCardList from '../UrlCardList';
import { useShortUrl } from '../../hooks/useShortUrl';
import { useGlobalStore } from '../../store/globalState';
import { useState } from 'react';
import CardListNav from '../CardListNav';
import { useErrorToast } from '../../hooks/useErrorToast';

const UserScreenContainer = () => {
	const [isRowsLayout, setIsRowsLayout] = useState(false);

	const { user, urls, error } = useGlobalStore();
	const { mutate: addUrl } = useShortUrl();
	useErrorToast(error);

	const onSetRowLayout = () => {
		setIsRowsLayout(true);
	};

	const onSetGridLayout = () => {
		setIsRowsLayout(false);
	};

	return (
		<main className='flex flex-col gap-5 h-full w-full items-center z-10'>
			<div className='flex flex-col gap-5 items-center'>
				<h2 className='text-[2rem]'>Short your favourite links!</h2>
				<ShortUrlForm
					addUrl={addUrl}
					userId={user?.id}
				/>
			</div>
			<CardListNav
				isRowsLayout={isRowsLayout}
				onSetRowLayout={onSetRowLayout}
				onSetGridLayout={onSetGridLayout}
			/>
			{urls && (
				<UrlCardList
					urls={urls}
					isRowsLayout={isRowsLayout}
				/>
			)}
		</main>
	);
};

export default UserScreenContainer;
