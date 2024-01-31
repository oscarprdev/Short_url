import { afterEach, beforeEach, describe, it } from 'vitest';
import { RenderResult, render } from '@testing-library/react';
import { mockedUrl } from '../../../tests/mocks/url-response';
import UrlCardCopyIcon from './UrlCardCopyIcon';

describe('UrlCardCopyIcon', () => {
	let component: RenderResult;

	beforeEach(() => {
		component = render(<UrlCardCopyIcon url={mockedUrl.shortUrl} />);
	});

	afterEach(() => {
		component.unmount();
	});

	it('Should render successfully', () => {
		component.getByTestId('icon-copy');
	});
});
