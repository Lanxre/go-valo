package client

func (h *httpClient) GetCroshair(options ...map[string]string) ([]byte, error) {
    data, err := h.GetRaw("/v1/crosshair/generate", options...)

	if err != nil {
        return nil, err
    }
    
    return data, err
}
