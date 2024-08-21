import RedirectingModal from './RedirectingModal';
import { RenderResult, render } from '@testing-library/react';
import { afterEach, beforeEach, describe, it } from 'vitest';

describe('RedirectingModal', () => {
	let component: RenderResult;

	beforeEach(() => {
		component = render(<RedirectingModal />);
	});

	afterEach(() => {
		component.unmount();
	});

	it('Should render successfully', () => {
		component.getByTestId('icon-loader');
		component.getByText('Redirecting...');
	});
});
