import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { RenderResult, fireEvent, render } from '@testing-library/react';
import CardListNav from './CardListNav';

describe('CardListNav', () => {
	let component: RenderResult;
	const onSetGridLayout = vi.fn();
	const onSetRowLayout = vi.fn();

	beforeEach(() => {
		component = render(
			<CardListNav
				isRowsLayout={true}
				onSetGridLayout={onSetGridLayout}
				onSetRowLayout={onSetRowLayout}
			/>
		);
	});

	afterEach(() => {
		component.unmount();
	});

	it('Should render successfully', async () => {
		component.getByLabelText('link-counter');
		component.getByTestId('icon-layout-list');
		component.getByTestId('icon-layout-grid');
	});

	it('Should trigger click events properly', () => {
		const listIcon = component.getByTestId('icon-layout-list');

		fireEvent.click(listIcon);

		expect(onSetRowLayout).toHaveBeenCalled();

		const gridIcon = component.getByTestId('icon-layout-grid');

		fireEvent.click(gridIcon);

		expect(onSetGridLayout).toHaveBeenCalled();
	});
});
