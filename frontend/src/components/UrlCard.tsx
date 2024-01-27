import { Url } from '../types/url';
import { formatTimeStamp } from '../utils/formatTimeStamp';
import UrlCardCopyIcon from './UrlCardCopyIcon';

interface UrlCardProps {
	url: Url;
	cardsRef: React.MutableRefObject<HTMLLIElement[]>;
	isHome?: boolean;
}

const UrlCard = ({ url, cardsRef, isHome }: UrlCardProps) => {
	return (
		<li
			key={url.id}
			ref={(el: HTMLLIElement | null) => el && cardsRef.current.push(el!)}
			className='relative flex flex-col gap-4 link-card w-fit h-fit border border-stone-800 font-light text-stone-300'>
			<a
				href={url.originalUrl}
				target='blank'
				className='hover:underline'>
				{url.shortUrl}
			</a>
			{!isHome && (
				<>
					<UrlCardCopyIcon url={url.shortUrl} />
					<div className='flex flex-col text-sm text-stone-400'>
						<p>Times used: {url.usage}</p>
						<p>Expiration: {formatTimeStamp(url.expiresAt)}</p>
					</div>
				</>
			)}
			<div className='glows'></div>
		</li>
	);
};

export default UrlCard;
