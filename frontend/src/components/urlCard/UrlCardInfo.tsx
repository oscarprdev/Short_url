import { formatTimeStamp } from '../../utils/formatTimeStamp';
import { IconPointer, IconCalendarX } from '@tabler/icons-react';
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from '../ui/tooltip';

interface UrlCardInfoProps {
	usage: number;
	expiresAt: string;
}

const UrlCardInfo = ({ usage, expiresAt }: UrlCardInfoProps) => {
	return (
		<div className='flex flex-col text-sm text-stone-400 w-full'>
			<div className='w-fit'>
				<TooltipProvider>
					<Tooltip>
						<TooltipTrigger>
							<div className='flex items-center gap-2'>
								<IconPointer className='w-4 text-stone-400' />
								<p>{usage}</p>
							</div>
						</TooltipTrigger>
						<TooltipContent>Times used</TooltipContent>
					</Tooltip>
				</TooltipProvider>
			</div>
			<div className='w-fit'>
				<TooltipProvider>
					<Tooltip>
						<TooltipTrigger>
							<div className='flex items-center gap-2'>
								<IconCalendarX className='w-4 text-stone-400' />
								<p>{formatTimeStamp(expiresAt)}</p>
							</div>
						</TooltipTrigger>
						<TooltipContent>Expiration date</TooltipContent>
					</Tooltip>
				</TooltipProvider>
			</div>
		</div>
	);
};

export default UrlCardInfo;
