import UrlCardInfo from './UrlCardInfo';
import { RenderResult, render } from '@testing-library/react';
import { afterEach, beforeEach, describe, it } from 'vitest';

describe('UrlCardInfo', () => {
	let component: RenderResult;
	const usage = 2;

	beforeEach(() => {
		component = render(<UrlCardInfo usage={usage} />);
	});

	afterEach(() => {
		component.unmount();
	});

	it('Should render successfully', async () => {
		component.getByTestId('icon-pointer');
		component.getByText(usage);
	});
});
