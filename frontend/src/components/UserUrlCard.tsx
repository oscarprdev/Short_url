import { Url } from '../types/url';
import { formatTimeStamp } from '../utils/formatTimeStamp';
import UrlCardCopyIcon from './UrlCardCopyIcon';

interface UserUrlCardProps {
	url: Url;
	cardsRef: React.MutableRefObject<HTMLLIElement[]>;
	isRowsLayout?: boolean;
}

const UserUrlCard = ({ url, cardsRef, isRowsLayout }: UserUrlCardProps) => {
	return (
		<li
			ref={(el: HTMLLIElement | null) => el && cardsRef.current.push(el!)}
			className={`${
				isRowsLayout && 'w-full'
			} relative flex flex-col gap-4 link-card py-5 h-fit border border-stone-800 font-light text-stone-300`}>
			<p>{url.titleUrl}</p>
			<a
				href={url.originalUrl}
				target='blank'
				className='hover:underline truncate'>
				{url.shortUrl}
			</a>
			<UrlCardCopyIcon url={url.shortUrl} />
			{!isRowsLayout && (
				<div className='flex flex-col text-sm text-stone-400'>
					<p>Times used: {url.usage}</p>
					<p>Expiration: {formatTimeStamp(url.expiresAt)}</p>
				</div>
			)}

			<div className='glows'></div>
		</li>
	);
};

export default UserUrlCard;
