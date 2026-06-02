package tests

import (
	ctrl "server/controllers"
	"server/models"
	"testing"
	"time"
)

func TestToLastUpdatedItemVariants(t *testing.T) {
	tests := map[string]struct {
		itemType string
		input    interface{}
		expected ctrl.LastUpdatedItem
	}{
		"reflection": {
			itemType: ctrl.ActivityReflection,
			input: models.ReflectionAnswer{
				Reflection: models.Reflection{
					Title: "Reflection Title",
					Phase: models.Phase{
						LifecycleID: 1,
						Lifecycle:   models.Lifecycle{Title: "Lifecycle Title"},
					},
				},
				UpdatedAt: time.Now(),
			},
			expected: ctrl.LastUpdatedItem{
				Type:           ctrl.ActivityReflection,
				Title:          "Reflection Title",
				LifecycleID:    1,
				LifecycleTitle: "Lifecycle Title",
			},
		},
		"further_reflection": {
			itemType: ctrl.ActivityFurtherReflection,
			input: models.FurtherReflectionAnswer{
				Reflection: models.Reflection{
					Title: "Further Reflection Title",
					Phase: models.Phase{
						LifecycleID: 2,
						Lifecycle:   models.Lifecycle{Title: "Lifecycle Title 2"},
					},
				},
				UpdatedAt: time.Now(),
			},
			expected: ctrl.LastUpdatedItem{
				Type:           ctrl.ActivityFurtherReflection,
				Title:          "Further Reflection Title",
				LifecycleID:    2,
				LifecycleTitle: "Lifecycle Title 2",
			},
		},
		"recommendation": {
			itemType: ctrl.ActivityRecommendation,
			input: models.RecommendationAnswer{
				Recommendation: models.Recommendation{
					ToolID: 3,
					Reflection: models.Reflection{
						Title: "Recommendation Reflection Title",
						Phase: models.Phase{
							LifecycleID: 4,
							Lifecycle:   models.Lifecycle{Title: "Lifecycle Title 4"},
						},
					},
				},
				UpdatedAt: time.Now(),
			},
			expected: ctrl.LastUpdatedItem{
				Type:           ctrl.ActivityRecommendation,
				Title:          "Recommendation Reflection Title",
				ToolID:         3,
				LifecycleID:    4,
				LifecycleTitle: "Lifecycle Title 4",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var item ctrl.LastUpdatedItem
			switch tc.itemType {
			case ctrl.ActivityReflection:
				item = ctrl.ToLastUpdatedItemReflection(tc.input.(models.ReflectionAnswer))
			case ctrl.ActivityFurtherReflection:
				item = ctrl.ToLastUpdatedItemFurtherReflection(tc.input.(models.FurtherReflectionAnswer))
			case ctrl.ActivityRecommendation:
				item = ctrl.ToLastUpdatedItemRecommendation(tc.input.(models.RecommendationAnswer))
			}
			if item.Type != tc.expected.Type {
				t.Errorf("expected type %s, got %s", tc.expected.Type, item.Type)
			}
			if item.Title != tc.expected.Title {
				t.Errorf("expected title %s, got %s", tc.expected.Title, item.Title)
			}
			if item.LifecycleID != tc.expected.LifecycleID {
				t.Errorf("expected lifecycleID %d, got %d", tc.expected.LifecycleID, item.LifecycleID)
			}
			if item.LifecycleTitle != tc.expected.LifecycleTitle {
				t.Errorf("expected lifecycleTitle %s, got %s", tc.expected.LifecycleTitle, item.LifecycleTitle)
			}
			if item.ToolID != tc.expected.ToolID {
				t.Errorf("expected toolID %d, got %d", tc.expected.ToolID, item.ToolID)
			}
		})
	}
}

func TestPickLatest(t *testing.T) {
	t1 := time.Now().Add(-1 * time.Hour)
	t2 := time.Now()
	item1 := ctrl.LastUpdatedItem{UpdatedAt: t1}
	item2 := ctrl.LastUpdatedItem{UpdatedAt: t2}
	latest := ctrl.PickLatest(nil, item1)
	if latest == nil || !latest.UpdatedAt.Equal(t1) {
		t.Errorf("expected latest to be item1")
	}
	latest = ctrl.PickLatest(latest, item2)
	if latest == nil || !latest.UpdatedAt.Equal(t2) {
		t.Errorf("expected latest to be item2")
	}
	latest = ctrl.PickLatest(latest, item1)
	if latest == nil || !latest.UpdatedAt.Equal(t2) {
		t.Errorf("expected latest to remain item2")
	}
}

func TestToLastUpdatedItem_EdgeCases(t *testing.T) {
	reflection := models.Reflection{Title: "", Phase: models.Phase{}}
	answer := models.ReflectionAnswer{Reflection: reflection, UpdatedAt: time.Time{}}
	item := ctrl.ToLastUpdatedItemReflection(answer)
	if item.Title != "" || item.LifecycleID != 0 || item.LifecycleTitle != "" {
		t.Errorf("expected zero values for missing phase/lifecycle, got %+v", item)
	}

	further := models.FurtherReflectionAnswer{Reflection: models.Reflection{}, UpdatedAt: time.Time{}}
	item2 := ctrl.ToLastUpdatedItemFurtherReflection(further)
	if item2.Title != "" || item2.LifecycleID != 0 || item2.LifecycleTitle != "" {
		t.Errorf("expected zero values for missing phase/lifecycle, got %+v", item2)
	}

	rec := models.RecommendationAnswer{Recommendation: models.Recommendation{}, UpdatedAt: time.Time{}}
	item3 := ctrl.ToLastUpdatedItemRecommendation(rec)
	if item3.Title != "" || item3.LifecycleID != 0 || item3.LifecycleTitle != "" || item3.ToolID != 0 {
		t.Errorf("expected zero values for missing recommendation/reflection, got %+v", item3)
	}
}

func TestToLastUpdatedItem_EmptyStringsAndZeroIDs(t *testing.T) {
	reflection := models.Reflection{Title: "", Phase: models.Phase{LifecycleID: 0, Lifecycle: models.Lifecycle{Title: ""}}}
	answer := models.ReflectionAnswer{Reflection: reflection, UpdatedAt: time.Now()}
	item := ctrl.ToLastUpdatedItemReflection(answer)
	if item.Title != "" || item.LifecycleID != 0 || item.LifecycleTitle != "" {
		t.Errorf("expected empty string and zero IDs, got %+v", item)
	}

	further := models.FurtherReflectionAnswer{Reflection: reflection, UpdatedAt: time.Now()}
	item2 := ctrl.ToLastUpdatedItemFurtherReflection(further)
	if item2.Title != "" || item2.LifecycleID != 0 || item2.LifecycleTitle != "" {
		t.Errorf("expected empty string and zero IDs, got %+v", item2)
	}

	rec := models.RecommendationAnswer{Recommendation: models.Recommendation{ToolID: 0, Reflection: reflection}, UpdatedAt: time.Now()}
	item3 := ctrl.ToLastUpdatedItemRecommendation(rec)
	if item3.Title != "" || item3.LifecycleID != 0 || item3.LifecycleTitle != "" || item3.ToolID != 0 {
		t.Errorf("expected empty string and zero IDs, got %+v", item3)
	}
}

func TestPickLatest_NilAndZero(t *testing.T) {
	var current *ctrl.LastUpdatedItem
	zero := ctrl.LastUpdatedItem{}
	result := ctrl.PickLatest(current, zero)
	if result == nil {
		t.Error("expected non-nil result for zero candidate")
	}
	old := ctrl.LastUpdatedItem{UpdatedAt: time.Now().Add(-2 * time.Hour)}
	cur := ctrl.LastUpdatedItem{UpdatedAt: time.Now()}
	res := ctrl.PickLatest(&cur, old)
	if res == nil || !res.UpdatedAt.Equal(cur.UpdatedAt) {
		t.Error("expected current to remain latest")
	}
}

func TestToLastUpdatedItem_NoPanicOnIncompleteData(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("mapping function panicked on incomplete data: %v", r)
		}
	}()
	_ = ctrl.ToLastUpdatedItemReflection(models.ReflectionAnswer{})
	_ = ctrl.ToLastUpdatedItemFurtherReflection(models.FurtherReflectionAnswer{})
	_ = ctrl.ToLastUpdatedItemRecommendation(models.RecommendationAnswer{})
}
