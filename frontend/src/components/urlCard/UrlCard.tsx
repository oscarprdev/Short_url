import { Url } from '../../types/url';

interface UrlCardProps {
	url: Url;
	cardsRef: React.MutableRefObject<HTMLLIElement[]>;
}

const UrlCard = ({ url, cardsRef }: UrlCardProps) => {
	return (
		<li
			ref={(el: HTMLLIElement | null) => el && cardsRef.current.push(el!)}
			className="relative flex flex-col gap-4 h-fit px-10 mt-5 py-5 rounded-lg border z-10 border-zinc-600 bg-contrast font-light text-black">
			<div className="w-full h-full p-2 absolute bg-white top-2 left-2 rounded-lg block">
				<a href={url.originalUrl} data-testid={`url-link-${url.id}`} target="blank" className="hover:underline">
					{url.shortUrl}
				</a>
			</div>
		</li>
	);
};

export default UrlCard;
