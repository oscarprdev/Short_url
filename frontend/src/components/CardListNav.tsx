import { IconLayoutGrid, IconLayoutList } from '@tabler/icons-react';

interface CardListNavProps {
	isRowsLayout: boolean;
	onSetRowLayout(): void;
	onSetGridLayout(): void;
}

const CardListNav = ({ isRowsLayout, onSetRowLayout, onSetGridLayout }: CardListNavProps) => {
	return (
		<nav className='w-[1100px] flex justify-center'>
			<div className='flex items-center gap-2'>
				<IconLayoutList
					onClick={onSetRowLayout}
					className={`${isRowsLayout ? 'text-stone-200' : 'text-stone-600 hover:text-stone-200'} cursor-pointer duration-300`}
				/>
				<IconLayoutGrid
					onClick={onSetGridLayout}
					className={`${!isRowsLayout ? 'text-stone-200' : 'text-stone-600 hover:text-stone-200'} cursor-pointer duration-300`}
				/>
			</div>
		</nav>
	);
};

export default CardListNav;
