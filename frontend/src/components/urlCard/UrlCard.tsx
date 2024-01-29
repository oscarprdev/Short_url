import { Url } from '../../types/url';

interface UrlCardProps {
	url: Url;
	cardsRef: React.MutableRefObject<HTMLLIElement[]>;
}

const UrlCard = ({ url, cardsRef }: UrlCardProps) => {
	return (
		<li
			ref={(el: HTMLLIElement | null) => el && cardsRef.current.push(el!)}
			className='relative flex flex-col gap-4 link-card h-fit px-10 py-5 border border-stone-800 font-light text-stone-300'>
			<a
				href={url.originalUrl}
				target='blank'
				className='hover:underline'>
				{url.shortUrl}
			</a>
			<div className='glows'></div>
		</li>
	);
};

export default UrlCard;
