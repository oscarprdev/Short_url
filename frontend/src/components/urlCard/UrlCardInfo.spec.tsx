import { afterEach, beforeEach, describe, it } from 'vitest';
import { RenderResult, render } from '@testing-library/react';
import UrlCardInfo from './UrlCardInfo';
import { formatTimeStamp } from '../../utils/formatTimeStamp';

describe('UrlCardInfo', () => {
	let component: RenderResult;
	const usage = 2;
	const expiresAt = '2022-01-31T23:59:59Z';

	beforeEach(() => {
		component = render(
			<UrlCardInfo
				usage={usage}
				expiresAt={expiresAt}
			/>
		);
	});

	afterEach(() => {
		component.unmount();
	});

	it('Should render successfully', async () => {
		component.getByTestId('icon-pointer');
		component.getByText(usage);

		component.getByTestId('icon-calendar');
		component.getByText(formatTimeStamp(expiresAt));
	});
});
