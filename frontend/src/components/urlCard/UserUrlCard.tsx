import { useUrlUsage } from '../../hooks/useUrlUsage';
import { Url } from '../../types/url';
import WrapperAction from '../containers/WrapperAction';
import UrlCardCopyIcon from './UrlCardCopyIcon';
import UrlCardInfo from './UrlCardInfo';
import UrlCardLink from './UrlCardLink';
import UrlCardTitle from './UrlCardTitle';

interface UserUrlCardProps {
	url: Url;
}

const UserUrlCard = ({ url }: UserUrlCardProps) => {
	const { mutate: urlUsage } = useUrlUsage();

	const onUrlClick = () => {
		urlUsage({ urlId: url.id });
	};

	return (
		<WrapperAction type="card" color="contrast" id="user-card">
			<UrlCardTitle title={url.titleUrl} urlId={url.id} />
			<UrlCardLink originalUrl={url.originalUrl} shortUrl={url.shortUrl} onUrlClick={onUrlClick} />
			<UrlCardCopyIcon url={url.shortUrl} />
			<UrlCardInfo usage={url.usage} />
		</WrapperAction>
	);
};

export default UserUrlCard;
