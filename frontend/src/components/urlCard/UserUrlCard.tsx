import { useDeleteUrl } from '../../hooks/useDeleteUrl';
import { useUrlUsage } from '../../hooks/useUrlUsage';
import { Url } from '../../types/url';
import WrapperAction from '../containers/WrapperAction';
import UrlCardCopyIcon from './UrlCardCopyIcon';
import UrlCardDeleteIcon from './UrlCardDeleteIcon';
import UrlCardInfo from './UrlCardInfo';
import UrlCardLink from './UrlCardLink';
import UrlCardTitle from './UrlCardTitle';

interface UserUrlCardProps {
	url: Url;
}

const UserUrlCard = ({ url }: UserUrlCardProps) => {
	const { mutate: urlUsage } = useUrlUsage();
	const { mutate: deleteUrl, isPending } = useDeleteUrl();

	const onUrlClick = () => {
		urlUsage({ urlId: url.id });
	};

	const onDeleteUrl = () => {
		deleteUrl({ urlId: url.id });
	};

	return (
		<WrapperAction type="card" color="contrast" id="user-card">
			<UrlCardTitle title={url.titleUrl} urlId={url.id} />
			<UrlCardLink originalUrl={url.originalUrl} shortUrl={url.shortUrl} onUrlClick={onUrlClick} />
			<UrlCardCopyIcon url={url.shortUrl} />
			<UrlCardDeleteIcon onDeleteUrl={onDeleteUrl} isPending={isPending}/>
			<UrlCardInfo usage={url.usage} />
		</WrapperAction>
	);
};

export default UserUrlCard;
