package services

import (
	"github.com/poseisharp/khairul-bot/internal/domain/entities"
	"github.com/poseisharp/khairul-bot/internal/persistent/repositories"
)

type PresetService struct {
	presetRepository *repositories.PresetRepository
}

func NewPresetService(presetRepository *repositories.PresetRepository) *PresetService {
	return &PresetService{
		presetRepository: presetRepository,
	}
}

func (s *PresetService) GetPreset(id int) (*entities.Preset, error) {
	preset, err := s.presetRepository.FindOne(id)
	if err != nil {
		return nil, err
	}

	return preset, nil
}

func (s *PresetService) CreatePreset(preset entities.Preset) error {
	return s.presetRepository.Store(preset)
}

func (s *PresetService) UpdatePreset(preset entities.Preset) error {
	return s.presetRepository.Update(preset)
}

func (s *PresetService) DeletePreset(id int) error {
	return s.presetRepository.Delete(id)
}

func (s *PresetService) GetPresetsByServerID(serverID string) ([]entities.Preset, error) {
	presets, err := s.presetRepository.FindByServerID(serverID)
	if err != nil {
		return nil, err
	}

	return presets, nil
}

func (s *PresetService) GetPresetByServerIDAndName(serverID string, name string) (*entities.Preset, error) {
	preset, err := s.presetRepository.FindByServerIDAndName(serverID, name)
	if err != nil {
		return nil, err
	}

	return preset, nil
}
