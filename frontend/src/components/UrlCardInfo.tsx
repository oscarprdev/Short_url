import { formatTimeStamp } from '../utils/formatTimeStamp';

interface UrlCardInfoProps {
	usage: number;
	expiresAt: string;
}

const UrlCardInfo = ({ usage, expiresAt }: UrlCardInfoProps) => {
	return (
		<div className='flex flex-col text-sm text-stone-400'>
			<p>Times used: {usage}</p>
			<p>Expiration: {formatTimeStamp(expiresAt)}</p>
		</div>
	);
};

export default UrlCardInfo;
