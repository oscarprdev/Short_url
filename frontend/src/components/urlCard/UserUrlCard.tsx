import { useUrlUsage } from '../../hooks/useUrlUsage';
import { Url } from '../../types/url';
import UrlCardCopyIcon from './UrlCardCopyIcon';
import UrlCardInfo from './UrlCardInfo';
import UrlCardLink from './UrlCardLink';
import UrlCardTitle from './UrlCardTitle';

interface UserUrlCardProps {
	url: Url;
	cardsRef: React.MutableRefObject<HTMLLIElement[]>;
	isRowsLayout?: boolean;
}

const UserUrlCard = ({ url, cardsRef, isRowsLayout }: UserUrlCardProps) => {
	const { mutate: urlUsage } = useUrlUsage();

	const onUrlClick = () => {
		urlUsage({ urlId: url.id });
	};

	return (
		<li
			ref={(el: HTMLLIElement | null) => el && cardsRef.current.push(el!)}
			className={`${
				isRowsLayout ? 'w-full' : 'w-[90vw] md:w-[30%]'
			} relative  flex flex-col gap-4 link-card py-5 h-fit border border-stone-800 font-light text-stone-300`}>
			<UrlCardTitle
				title={url.titleUrl}
				urlId={url.id}
			/>
			<UrlCardLink
				originalUrl={url.originalUrl}
				shortUrl={url.shortUrl}
				onUrlClick={onUrlClick}
			/>
			<UrlCardCopyIcon url={url.shortUrl} />
			{!isRowsLayout && (
				<UrlCardInfo
					usage={url.usage}
					expiresAt={url.expiresAt}
				/>
			)}
			<div className='glows'></div>
		</li>
	);
};

export default UserUrlCard;
