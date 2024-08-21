import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from '../ui/tooltip';
import { IconPointer } from '@tabler/icons-react';

interface UrlCardInfoProps {
	usage: number;
}

const UrlCardInfo = ({ usage }: UrlCardInfoProps) => {
	return (
		<div className="flex flex-col text-sm text-stone-400 w-full">
			<div className="w-fit">
				<TooltipProvider>
					<Tooltip>
						<TooltipTrigger>
							<div className="flex items-center gap-2 w-[170px] text-xs">
								<IconPointer data-testid="icon-pointer" size={14} className="text-stone-400" />
								{usage}
							</div>
						</TooltipTrigger>
						<TooltipContent>Times used</TooltipContent>
					</Tooltip>
				</TooltipProvider>
			</div>
		</div>
	);
};

export default UrlCardInfo;
