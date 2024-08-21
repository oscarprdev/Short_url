import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from '../ui/tooltip';
import { IconLoader2, IconTrash } from '@tabler/icons-react';

type UrlCardDeleteIconProps = {
	onDeleteUrl(): void;
	isPending: boolean;
};

const UrlCardDeleteIcon = ({ onDeleteUrl, isPending }: UrlCardDeleteIconProps) => {
	return (
		<span className="group-hover:block absolute hidden top-2 right-7">
			<TooltipProvider>
				<Tooltip>
					<TooltipTrigger>
						{isPending ? (
							<IconLoader2 className="animate-spin w-4 text-stone-300 hover:text-stone-100" />
						) : (
							<IconTrash
								data-testid="icon-copy"
								onClick={onDeleteUrl}
								className="animate-fade-in w-4 text-stone-300 hover:text-stone-100 duration-200"
							/>
						)}
					</TooltipTrigger>
					<TooltipContent>Delete link</TooltipContent>
				</Tooltip>
			</TooltipProvider>
		</span>
	);
};

export default UrlCardDeleteIcon;
