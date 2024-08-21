import { Url } from '../../types/url';
import WrapperAction from '../containers/WrapperAction';

interface UrlCardProps {
	url: Url;
}

const UrlCard = ({ url }: UrlCardProps) => {
	return (
		<WrapperAction size="md" type="url" color="contrast" url={url.originalUrl} id={url.id}>
			{url.shortUrl}
		</WrapperAction>
	);
};

export default UrlCard;
