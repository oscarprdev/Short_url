interface UrlCardLinkProps {
	originalUrl: string;
	shortUrl: string;
	onUrlClick(): void;
}
const UrlCardLink = ({ originalUrl, shortUrl, onUrlClick }: UrlCardLinkProps) => {
	return (
		<a
			href={originalUrl}
			onClick={onUrlClick}
			target='blank'
			className='hover:underline truncate'>
			{shortUrl}
		</a>
	);
};

export default UrlCardLink;
